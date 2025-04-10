package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/esc-chula/intania-vote/apps/api/zk"
)

type BallotService interface {
	CreateBallotProof(ctx context.Context, oidcId string, choiceId uint) (*zk.Proof, error)
	CreateBallot(ctx context.Context, oidcId string, voteSlug string, proof zk.Proof) (*model.Ballot, *string, error)
}

type ballotServiceImpl struct {
	ballotRepo repository.BallotRepository
	userRepo   repository.UserRepository
	voteRepo   repository.VoteRepository
	cfg        *config.Config
}

var _ BallotService = &ballotServiceImpl{}

func NewBallotService(ballotRepo repository.BallotRepository, userRepo repository.UserRepository, voteRepo repository.VoteRepository, cfg *config.Config) BallotService {
	return &ballotServiceImpl{
		ballotRepo: ballotRepo,
		userRepo:   userRepo,
		voteRepo:   voteRepo,
		cfg:        cfg,
	}
}

func (s *ballotServiceImpl) CreateBallotProof(ctx context.Context, oidcId string, choiceId uint) (*zk.Proof, error) {
	secret := zk.DeriveSecretFromSub(oidcId, s.cfg)

	proofData, err := zk.GenerateProof(choiceId, secret)
	if err != nil {
		return nil, err
	}

	return proofData, nil
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

	isUserCreated, err := s.ballotRepo.IsUserCreated(ctx, user.Id, vote.Id)
	if err != nil {
		return nil, nil, err
	}
	if isUserCreated {
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

	isHasValidProof := false
	var choiceId uint
	for _, c := range vote.Choices {
		isValid := zk.VerifyProof(proofMap, c.Id, secret)

		if isValid {
			isHasValidProof = true
			choiceId = c.Id
			break
		}
	}
	if !isHasValidProof {
		return nil, nil, errors.New("invalid proof")
	}

	encryptedBallot := zk.EncryptWithServerKey(choiceId, s.cfg)

	encryptedKey := zk.GenerateVoteEncryptionKey(user.OidcId.String(), proof.Nullifier)

	encryptedChoiceId := zk.EncryptChoiceId(choiceId, encryptedKey)

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

	base58Receipt := base58.Encode([]byte(fmt.Sprintf("%d:%s:%s", vote.Id, receipt, encryptedChoiceId)))

	return ballot, &base58Receipt, nil
}
