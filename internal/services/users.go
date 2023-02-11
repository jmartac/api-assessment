package services

import (
	"api-assessment/internal/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	repositories.UserRepository
}

// NewUserService creates and returns a new UserService
func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		UserRepository: repositories.NewUserRepository(db),
	}
}

type userServiceImpl struct {
	repositories.UserRepository
}
