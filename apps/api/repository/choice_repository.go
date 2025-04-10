package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type ChoiceRepository interface {
	GetById(ctx context.Context, id uint) (*model.Choice, error)
}

type choiceRepositoryImpl struct {
	db *gorm.DB
}

var _ ChoiceRepository = &choiceRepositoryImpl{}

func NewChoiceRepository(db *gorm.DB) ChoiceRepository {
	return &choiceRepositoryImpl{
		db: db,
	}
}

func (r *choiceRepositoryImpl) GetById(ctx context.Context, id uint) (*model.Choice, error) {
	var choice model.Choice
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&choice).Error
	if err != nil {
		return nil, err
	}
	return &choice, nil
}
