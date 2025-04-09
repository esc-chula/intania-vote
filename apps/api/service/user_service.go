package service

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, oidcId string, studentId string) error
	GetUserByOidcIdOrStudentId(ctx context.Context, oidcId string, studentId string) (*model.User, error)
	UpdateUserById(ctx context.Context, id uint, oidcId string, studentId string) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

var _ UserService = &userServiceImpl{}

func NewUserService(userRepo repository.UserRepository, cfg *config.Config) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, oidcId string, studentId string) error {
	oidcIdUUID, err := uuid.Parse(oidcId)
	if err != nil {
		return err
	}

	user := &model.User{
		OidcId:    oidcIdUUID,
		StudentId: studentId,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) GetUserByOidcIdOrStudentId(ctx context.Context, oidcId string, studentId string) (*model.User, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		user, err = s.userRepo.GetByStudentId(ctx, studentId)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s *userServiceImpl) UpdateUserById(ctx context.Context, id uint, oidcId string, studentId string) error {
	user := &model.User{
		OidcId:    uuid.MustParse(oidcId),
		StudentId: studentId,
	}

	if err := s.userRepo.UpdateById(ctx, id, user); err != nil {
		return err
	}
	return nil
}
