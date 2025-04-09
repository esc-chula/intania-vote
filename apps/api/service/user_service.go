package service

import (
	"context"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, oidcId uuid.UUID, studentId string) error
	GetUserByOidcId(ctx context.Context, oidcId string) (*model.User, error)
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

func (s *userServiceImpl) CreateUser(ctx context.Context, oidcId uuid.UUID, studentId string) error {
	user := &model.User{
		OidcId:    oidcId,
		StudentId: studentId,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) GetUserByOidcId(ctx context.Context, oidcId string) (*model.User, error) {
	user, err := s.userRepo.GetByOidcId(ctx, oidcId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
