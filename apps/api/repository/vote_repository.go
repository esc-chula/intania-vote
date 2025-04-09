package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type VoteRepository interface {
	Create(ctx context.Context, vote *model.Vote, choices []*model.Choice) (*model.Vote, error)
	Get(ctx context.Context) ([]*model.Vote, error)
	GetById(ctx context.Context, id uint) (*model.Vote, error)
	GetBySlug(ctx context.Context, slug string) (*model.Vote, error)
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

func (v *voteRepositoryImpl) Get(ctx context.Context) ([]*model.Vote, error) {
	var votes []*model.Vote
	if err := v.db.WithContext(ctx).Preload("Choices").Find(&votes).Error; err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *voteRepositoryImpl) GetById(ctx context.Context, id uint) (*model.Vote, error) {
	var vote model.Vote
	if err := v.db.WithContext(ctx).Preload("Choices").First(&vote, id).Error; err != nil {
		return nil, err
	}
	return &vote, nil
}

func (v *voteRepositoryImpl) GetBySlug(ctx context.Context, slug string) (*model.Vote, error) {
	var vote model.Vote
	if err := v.db.WithContext(ctx).Preload("Choices").Where("slug = ?", slug).First(&vote).Error; err != nil {
		return nil, err
	}
	return &vote, nil
}
