package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title       string
	Director    string `gorm:"not_null"`
	ReleaseDate string `gorm:"not_null"`
	Genre       string `gorm:"not_null"`
	Synopsis    string `gorm:"not_null"`
	// TODO Cast
}
