package service

import (
	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/repository"
)

type UserService interface {
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
