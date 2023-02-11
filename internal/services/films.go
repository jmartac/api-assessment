package services

import (
	"api-assessment/internal/repositories"
	"gorm.io/gorm"
)

type FilmService interface {
	repositories.FilmRepository
}

// NewFilmService creates and returns a new FilmService
func NewFilmService(db *gorm.DB) FilmService {
	return &filmServiceImpl{
		FilmRepository: repositories.NewFilmRepository(db),
	}
}

type filmServiceImpl struct {
	repositories.FilmRepository
}
