package service

import (
	"context"
	"regexp"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/esc-chula/intania-vote/apps/api/zk"
	"gorm.io/gorm"
)

type VoteService interface {
	CreateVote(ctx context.Context, oidcId string, vote *model.Vote, choices []*model.Choice) (*model.Vote, error)
	GetVoteById(ctx context.Context, id uint) (*model.Vote, error)
	GetVoteBySlug(ctx context.Context, slug string) (*model.Vote, error)
	GetVotes(ctx context.Context) ([]*model.Vote, error)
	GetVotesByUserEligibility(ctx context.Context, oidcId string) ([]*model.Vote, error)
	TallyVoteById(ctx context.Context, id uint) (*model.Tally, error)
}

type voteServiceImpl struct {
	voteRepo   repository.VoteRepository
	userRepo   repository.UserRepository
	ballotRepo repository.BallotRepository
	cfg        *config.Config
}

var _ VoteService = &voteServiceImpl{}

func NewVoteService(voteRepo repository.VoteRepository, userRepo repository.UserRepository, ballotRepo repository.BallotRepository, cfg *config.Config) VoteService {
	return &voteServiceImpl{
		voteRepo:   voteRepo,
		userRepo:   userRepo,
		ballotRepo: ballotRepo,
		cfg:        cfg,
	}
}

func (s *voteServiceImpl) CreateVote(ctx context.Context, oidcId string, vote *model.Vote, choices []*model.Choice) (*model.Vote, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, err
	}

	vote.UserId = user.Id

	createdVote, err := s.voteRepo.Create(ctx, vote, choices)
	if err != nil {
		return nil, err
	}

	return createdVote, nil
}

func (s *voteServiceImpl) GetVoteById(ctx context.Context, id uint) (*model.Vote, error) {
	vote, err := s.voteRepo.GetById(ctx, id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil
	}

	return vote, nil
}

func (s *voteServiceImpl) GetVoteBySlug(ctx context.Context, slug string) (*model.Vote, error) {
	vote, err := s.voteRepo.GetBySlug(ctx, slug)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil
	}

	return vote, nil
}

func (s *voteServiceImpl) GetVotes(ctx context.Context) ([]*model.Vote, error) {
	votes, err := s.voteRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	if len(votes) == 0 {
		return nil, nil
	}

	return votes, nil
}

func (s *voteServiceImpl) GetVotesByUserEligibility(ctx context.Context, oidcId string) ([]*model.Vote, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, err
	}

	activeVotes, err := s.voteRepo.GetByActiveTime(ctx)
	if err != nil {
		return nil, err
	}
	if len(activeVotes) == 0 {
		return nil, nil
	}

	eligibleVotes := make([]*model.Vote, 0)
	for _, vote := range activeVotes {
		if vote.EligibleStudentId == "*" || (user.StudentId != "" && regexp.MustCompile(vote.EligibleStudentId).MatchString(user.StudentId)) {
			eligibleVotes = append(eligibleVotes, vote)
			continue
		}
	}

	if len(eligibleVotes) == 0 {
		return nil, nil
	}

	return eligibleVotes, nil
}

func (s *voteServiceImpl) TallyVoteById(ctx context.Context, id uint) (*model.Tally, error) {
	vote, err := s.voteRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if vote == nil {
		return nil, nil
	}

	choices := vote.Choices

	ballots, err := s.ballotRepo.GetBallotsByVoteId(ctx, id)
	if err != nil {
		return nil, err
	}

	choiceBallotCounts := make(map[int]int)

	for _, ballot := range ballots {
		choiceId, err := zk.DecryptWithServerKey(ballot.EncryptedBallot, s.cfg)
		if err != nil {
			continue
		}

		found := false
		for _, choice := range choices {
			if choice.Id == uint(choiceId) {
				choiceBallotCounts[choiceId]++
				found = true
				break
			}
		}

		if !found {
			if choiceId == 0 {
				choiceBallotCounts[0]++
			} else if choiceId == -1 {
				choiceBallotCounts[-1]++
			}
		}
	}

	var tally []model.TallyChoices
	totalBallots := 0

	for _, choice := range choices {
		count := choiceBallotCounts[int(choice.Id)]
		tally = append(tally, model.TallyChoices{
			Number: int(choice.Number),
			Count:  uint(count),
		})
		totalBallots += count
	}

	if noVotes := choiceBallotCounts[0]; noVotes > 0 {
		tally = append(tally, model.TallyChoices{
			Number: 0,
			Count:  uint(noVotes),
		})
		totalBallots += noVotes
	}

	if voteNos := choiceBallotCounts[-1]; voteNos > 0 {
		tally = append(tally, model.TallyChoices{
			Number: -1,
			Count:  uint(voteNos),
		})
		totalBallots += voteNos
	}

	return &model.Tally{
		Choices: tally,
		Total:   uint(totalBallots),
	}, nil
}
