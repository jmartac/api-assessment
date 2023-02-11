package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not_null;unique" json:"username"`
	Password     string `gorm:"-" json:"password"`
	PasswordHash string `gorm:"not_null" json:"-"`
}

type UserService interface {
	UserRepository
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
	UsernameExists(username string) (bool, error)
	Create(user *User) error
}

// NewUserService creates and returns a new UserService
func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{
		UserRepository: &userRepositoryImpl{
			db: db,
		},
	}
}

type userServiceImpl struct {
	UserRepository
}

type userRepositoryImpl struct {
	db *gorm.DB
}

// FindByUsername returns the first user with the given ID
func (repoImpl *userRepositoryImpl) FindByUsername(username string) (*User, error) {
	var user User
	err := repoImpl.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UsernameExists returns true if the username is already in use
func (repoImpl *userRepositoryImpl) UsernameExists(username string) (bool, error) {
	var user User
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
func (repoImpl *userRepositoryImpl) Create(user *User) error {
	err := repoImpl.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
