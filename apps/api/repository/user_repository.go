package repository

import "gorm.io/gorm"

type UserRepository interface {
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
