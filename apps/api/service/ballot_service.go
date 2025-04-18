package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/esc-chula/intania-vote/apps/api/zk"
)

type BallotService interface {
	CreateBallotProof(ctx context.Context, oidcId string, voteSlug string, choiceNumber string) (*zk.Proof, error)
	CreateBallot(ctx context.Context, oidcId string, voteSlug string, proof zk.Proof) (*model.Ballot, *string, error)
	VerifyBallot(ctx context.Context, oidcId string, ballotKey string) (*int, *time.Time, error)
	// TallyByVoteSlug(ctx context.Context, voteSlug string) ([]*model.Tally, error)
	GetBallotsByOidcId(ctx context.Context, oidcId string) ([]*model.Ballot, error)
	HasUserVoted(ctx context.Context, oidcId string, voteSlug string) (bool, error)
	CountBallotsByVoteId(ctx context.Context, voteId uint) (int64, error)
}

type ballotServiceImpl struct {
	ballotRepo repository.BallotRepository
	userRepo   repository.UserRepository
	voteRepo   repository.VoteRepository
	choiceRepo repository.ChoiceRepository
	cfg        *config.Config
}

var _ BallotService = &ballotServiceImpl{}

func NewBallotService(ballotRepo repository.BallotRepository, userRepo repository.UserRepository, voteRepo repository.VoteRepository, choiceRepo repository.ChoiceRepository, cfg *config.Config) BallotService {
	return &ballotServiceImpl{
		ballotRepo: ballotRepo,
		userRepo:   userRepo,
		voteRepo:   voteRepo,
		choiceRepo: choiceRepo,
		cfg:        cfg,
	}
}

func (s *ballotServiceImpl) CreateBallotProof(ctx context.Context, oidcId string, voteSlug string, choiceNumber string) (*zk.Proof, error) {
	secret := zk.DeriveSecretFromSub(oidcId, s.cfg)

	vote, err := s.voteRepo.GetBySlug(ctx, voteSlug)
	if err != nil {
		return nil, err
	}

	switch choiceNumber {
	case "0":
		proofData, err := zk.GenerateProof(0, secret)
		if err != nil {
			return nil, err
		}

		return proofData, nil
	case "-1":
		proofData, err := zk.GenerateProof(-1, secret)
		if err != nil {
			return nil, err
		}

		return proofData, nil
	default:
		choice, err := s.choiceRepo.GetByVoteIdAndNumber(ctx, vote.Id, choiceNumber)
		if err != nil {
			return nil, err
		}

		proofData, err := zk.GenerateProof(int(choice.Id), secret)
		if err != nil {
			return nil, err
		}

		return proofData, nil
	}
}

func (s *ballotServiceImpl) CreateBallot(ctx context.Context, oidcId string, voteSlug string, proof zk.Proof) (*model.Ballot, *string, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, nil, err
	}

	vote, err := s.voteRepo.GetBySlug(ctx, voteSlug)
	if err != nil {
		return nil, nil, err
	}
	if vote == nil {
		return nil, nil, errors.New("vote not found")
	}

	hasVoted, err := s.ballotRepo.HasUserVoted(ctx, user.Id, vote.Id)
	if err != nil {
		return nil, nil, err
	}
	if hasVoted {
		return nil, nil, errors.New("user already created ballot for this vote")
	}

	secret := zk.DeriveSecretFromSub(user.OidcId.String(), s.cfg)

	proofMap := map[string]string{
		"commitment":     proof.Commitment,
		"blindingFactor": proof.BlindingFactor,
		"challenge":      proof.Challenge,
		"response":       proof.Response,
		"nullifier":      proof.Nullifier,
	}

	// try 0 and -1 first
	// if not found, try the choices
	isHasValidProof := false
	var choiceId int

	isNoVoteValid := zk.VerifyProof(proofMap, 0, secret)
	isVoteNotValid := zk.VerifyProof(proofMap, -1, secret)
	if isNoVoteValid {
		isHasValidProof = true
		choiceId = 0
	} else if isVoteNotValid {
		isHasValidProof = true
		choiceId = -1
	} else {
		for _, c := range vote.Choices {
			isValid := zk.VerifyProof(proofMap, int(c.Id), secret)

			if isValid {
				isHasValidProof = true
				choiceId = int(c.Id)
				break
			}
		}
	}
	if !isHasValidProof {
		return nil, nil, errors.New("invalid proof")
	}

	encryptedBallot := zk.EncryptWithServerKey(choiceId, s.cfg)

	encryptionKey := zk.GenerateVoteEncryptionKey(user.OidcId.String(), proof.Nullifier)

	encryptedChoiceId := zk.EncryptChoiceId(choiceId, encryptionKey)

	receiptData := fmt.Sprintf("%s:%s:%s",
		proof.Commitment,
		proof.Nullifier,
		encryptedChoiceId)

	receiptHash := sha256.Sum256([]byte(receiptData))
	receipt := base64.StdEncoding.EncodeToString(receiptHash[:])

	ballot, err := s.ballotRepo.Create(ctx, &model.Ballot{
		UserId:          user.Id,
		VoteId:          vote.Id,
		Challenge:       proof.Challenge,
		Commitment:      proof.Commitment,
		BlindingFactor:  proof.BlindingFactor,
		Response:        proof.Response,
		EncryptedBallot: encryptedBallot,
		Nullifier:       proof.Nullifier,
	})
	if err != nil {
		return nil, nil, err
	}

	if choiceId > 0 {
		err = s.choiceRepo.IncrementVoteCountById(ctx, uint(choiceId))
		if err != nil {
			return nil, nil, err
		}
	}

	base58Receipt := base58.Encode([]byte(fmt.Sprintf("%d:%s:%s", vote.Id, receipt, encryptedChoiceId)))

	return ballot, &base58Receipt, nil
}

func (s *ballotServiceImpl) VerifyBallot(ctx context.Context, oidcId string, ballotKey string) (*int, *time.Time, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, nil, err
	}

	decodedBallotKey := base58.Decode(ballotKey)
	if decodedBallotKey == nil {
		return nil, nil, errors.New("invalid ballot key")
	}
	ballotKeyString := string(decodedBallotKey) // voteId:receipt:encryptedChoiceId
	parts := strings.Split(ballotKeyString, ":")
	if len(parts) != 3 {
		return nil, nil, errors.New("invalid ballot key format")
	}
	voteId := parts[0]
	receipt := parts[1]
	encryptedChoiceId := parts[2]
	if voteId == "" || receipt == "" || encryptedChoiceId == "" {
		return nil, nil, errors.New("invalid ballot key format")
	}

	voteIdUint, err := strconv.ParseUint(voteId, 10, 32)
	if err != nil {
		log.Println(1)
		return nil, nil, err
	}
	if voteIdUint == 0 {
		return nil, nil, errors.New("invalid vote id")
	}

	ballots, err := s.ballotRepo.GetBallotsByUserId(ctx, user.Id)
	if err != nil {
		log.Println(2)
		return nil, nil, err
	}

	for _, ballot := range ballots {
		encryptionKey := zk.GenerateVoteEncryptionKey(user.OidcId.String(), ballot.Nullifier)

		choiceId, err := zk.DecryptChoiceId(encryptedChoiceId, encryptionKey)
		if err != nil {
			continue
		}

		receiptData := fmt.Sprintf("%s:%s:%s",
			ballot.Commitment,
			ballot.Nullifier,
			encryptedChoiceId)

		receiptHash := sha256.Sum256([]byte(receiptData))
		expectedReceipt := base64.StdEncoding.EncodeToString(receiptHash[:])

		if expectedReceipt == receipt {
			var choice int

			if choiceId <= 0 {
				choice = choiceId
			} else {
				choiceData, err := s.choiceRepo.GetById(ctx, uint(choiceId))
				if err != nil {
					log.Println(4)
					return nil, nil, err
				}
				choice = int(choiceData.Number)
			}

			return &choice, &ballot.CreatedAt, nil
		}
	}

	return nil, nil, errors.New("ballot not found")
}

func (s *ballotServiceImpl) GetBallotsByOidcId(ctx context.Context, oidcId string) ([]*model.Ballot, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, err
	}

	ballots, err := s.ballotRepo.GetBallotsByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return ballots, nil
}

func (s *ballotServiceImpl) HasUserVoted(ctx context.Context, oidcId string, voteSlug string) (bool, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, errors.New("user not found")
	}

	vote, err := s.voteRepo.GetBySlug(ctx, voteSlug)
	if err != nil {
		return false, err
	}
	if vote == nil {
		return false, errors.New("vote not found")
	}

	hasVoted, err := s.ballotRepo.HasUserVoted(ctx, user.Id, vote.Id)
	if err != nil {
		return false, err
	}
	return hasVoted, nil
}

func (s *ballotServiceImpl) CountBallotsByVoteId(ctx context.Context, voteId uint) (int64, error) {
	count, err := s.ballotRepo.CountByVoteId(ctx, voteId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
