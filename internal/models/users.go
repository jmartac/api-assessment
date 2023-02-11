package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not_null;unique" json:"username"`
	Password     string `gorm:"-" json:"password"`
	PasswordHash string `gorm:"not_null" json:"-"`
}
