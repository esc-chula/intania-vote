package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type BallotRepository interface {
	Create(ctx context.Context, ballot *model.Ballot) (*model.Ballot, error)
	GetBallotsByUserId(ctx context.Context, userId uint) ([]*model.Ballot, error)
	GetBallotsByVoteId(ctx context.Context, voteId uint) ([]*model.Ballot, error)
	HasUserVoted(ctx context.Context, userId uint, voteId uint) (bool, error)
	CountByVoteId(ctx context.Context, voteId uint) (int64, error)
}

type ballotRepositoryImpl struct {
	db *gorm.DB
}

var _ BallotRepository = &ballotRepositoryImpl{}

func NewBallotRepository(db *gorm.DB) BallotRepository {
	return &ballotRepositoryImpl{
		db: db,
	}
}

func (ballotRepo *ballotRepositoryImpl) Create(ctx context.Context, ballot *model.Ballot) (*model.Ballot, error) {
	if err := ballotRepo.db.WithContext(ctx).Create(ballot).Error; err != nil {
		return nil, err
	}
	return ballot, nil
}

func (ballotRepo *ballotRepositoryImpl) GetBallotsByUserId(ctx context.Context, userId uint) ([]*model.Ballot, error) {
	var ballots []*model.Ballot
	if err := ballotRepo.db.WithContext(ctx).Where("user_id = ?", userId).Find(&ballots).Error; err != nil {
		return nil, err
	}

	return ballots, nil
}

func (ballotRepo *ballotRepositoryImpl) GetBallotsByVoteId(ctx context.Context, voteId uint) ([]*model.Ballot, error) {
	var ballots []*model.Ballot
	if err := ballotRepo.db.WithContext(ctx).Where("vote_id = ?", voteId).Find(&ballots).Error; err != nil {
		return nil, err
	}

	return ballots, nil
}

func (ballotRepo *ballotRepositoryImpl) HasUserVoted(ctx context.Context, userId uint, voteId uint) (bool, error) {
	var count int64
	if err := ballotRepo.db.WithContext(ctx).Model(&model.Ballot{}).
		Where("user_id = ? AND vote_id = ?", userId, voteId).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ballotRepo *ballotRepositoryImpl) CountByVoteId(ctx context.Context, voteId uint) (int64, error) {
	var count int64
	if err := ballotRepo.db.WithContext(ctx).Model(&model.Ballot{}).
		Where("vote_id = ?", voteId).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
