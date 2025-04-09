package service

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"gorm.io/gorm"
)

type VoteService interface {
	CreateVote(ctx context.Context, oidcId string, vote *model.Vote, choices []*model.Choice) (*model.Vote, error)
	GetVoteById(ctx context.Context, id uint) (*model.Vote, error)
	GetVoteBySlug(ctx context.Context, slug string) (*model.Vote, error)
	GetUserEligibleVotes(ctx context.Context, userId uint) ([]*model.Vote, error)
}

type voteServiceImpl struct {
	voteRepo repository.VoteRepository
	userRepo repository.UserRepository
	cfg      *config.Config
}

var _ VoteService = &voteServiceImpl{}

func NewVoteService(voteRepo repository.VoteRepository, userRepo repository.UserRepository, cfg *config.Config) VoteService {
	return &voteServiceImpl{
		voteRepo: voteRepo,
		userRepo: userRepo,
		cfg:      cfg,
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

func (s *voteServiceImpl) GetUserEligibleVotes(ctx context.Context, userId uint) ([]*model.Vote, error) {
	return nil, nil
}
