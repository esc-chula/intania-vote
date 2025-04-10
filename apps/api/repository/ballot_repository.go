package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type BallotRepository interface {
	Create(ctx context.Context, ballot *model.Ballot) (*model.Ballot, error)
	IsUserCreated(ctx context.Context, userId uint, voteId uint) (bool, error)
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

func (ballotRepo *ballotRepositoryImpl) IsUserCreated(ctx context.Context, userId uint, voteId uint) (bool, error) {
	var count int64
	if err := ballotRepo.db.WithContext(ctx).Model(&model.Ballot{}).
		Where("user_id = ? AND vote_id = ?", userId, voteId).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
