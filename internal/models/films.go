package models

import (
	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Title       string `gorm:"not_null;unique" json:"title"`
	Director    string `gorm:"not_null" json:"director"`
	ReleaseDate string `gorm:"not_null" json:"release_date"`
	Genre       string `gorm:"not_null" json:"genre"`
	Synopsis    string `gorm:"not_null" json:"synopsis"`
	// TODO Cast
}
