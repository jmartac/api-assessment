package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not_null;unique"`
	PasswordHash string `gorm:"not_null"`
	Films        []Film
}

// ToResponse converts a User model into a UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Films:    Films(u.Films).ToResponse(),
	}
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Films    []FilmResponse `json:"films,omitempty"`
}
