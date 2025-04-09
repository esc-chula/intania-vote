package repository

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, id uint) (*model.User, error)
	GetByOidcId(ctx context.Context, oidcId string) (*model.User, error)
	GetByStudentId(ctx context.Context, studentId string) (*model.User, error)
	UpdateById(ctx context.Context, id uint, user *model.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

var _ UserRepository = &userRepositoryImpl{}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepositoryImpl) GetById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByOidcId(ctx context.Context, oidcId string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("oidc_id = ?", oidcId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByStudentId(ctx context.Context, studentId string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("student_id = ?", studentId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateById(ctx context.Context, id uint, user *model.User) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(user).Error
}
