package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title       string `gorm:"not_null" json:"title"`
	Director    string `gorm:"not_null" json:"director"`
	ReleaseDate string `gorm:"not_null" json:"release_date"`
	Genre       string `gorm:"not_null" json:"genre"`
	Synopsis    string `gorm:"not_null" json:"synopsis"`
	// TODO Cast
}

type FilmService interface {
	FilmRepository
}

type FilmRepository interface {
	FindAll() ([]Film, error)
}

// NewFilmService creates and returns a new FilmService
func NewFilmService(db *gorm.DB) FilmService {
	return &filmServiceImpl{
		FilmRepository: &filmRepositoryImpl{
			db: db,
		},
	}
}

type filmServiceImpl struct {
	FilmRepository
}

type filmRepositoryImpl struct {
	db *gorm.DB
}

func (repoImpl *filmRepositoryImpl) FindAll() ([]Film, error) {
	var films []Film
	err := repoImpl.db.Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}
