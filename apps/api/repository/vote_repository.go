package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type VoteRepository interface {
	Create(ctx context.Context, vote *model.Vote, choices []*model.Choice) (*model.Vote, error)
}

type voteRepositoryImpl struct {
	db *gorm.DB
}

var _ VoteRepository = &voteRepositoryImpl{}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &voteRepositoryImpl{
		db: db,
	}
}

func (v *voteRepositoryImpl) Create(ctx context.Context, vote *model.Vote, choices []*model.Choice) (*model.Vote, error) {
	tx := v.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(vote).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := range choices {
		choices[i].VoteId = vote.Id
	}

	if err := tx.Model(vote).Association("Choices").Append(choices); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := v.db.Preload("Choices").First(vote, vote.Id).Error; err != nil {
		return nil, err
	}

	return vote, nil
}
