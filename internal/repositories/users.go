package repositories

import (
	"api-assessment/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	UsernameExists(username string) (bool, error)
	Create(user *models.User) error
}

// NewUserRepository creates and returns a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

type userRepositoryImpl struct {
	db *gorm.DB
}

// FindByUsername returns the first user with the given ID
func (repoImpl *userRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repoImpl.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UsernameExists returns true if the username is already in use
func (repoImpl *userRepositoryImpl) UsernameExists(username string) (bool, error) {
	var user models.User
	err := repoImpl.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Create will insert a given user in the DB
func (repoImpl *userRepositoryImpl) Create(user *models.User) error {
	err := repoImpl.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
