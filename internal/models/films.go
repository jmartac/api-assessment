package models

import (
	"gorm.io/gorm"
	"strings"
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

type FilmService interface {
	FilmRepository
}

type FilmRepository interface {
	FindAll(title, genre, releaseDate string) ([]Film, error)
	FindByID(id uint) (*Film, error)
	Create(film *Film) error
	Update(film *Film) error
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

// FindAll returns all films in the database
func (repoImpl *filmRepositoryImpl) FindAll(title, genre, releaseDate string) ([]Film, error) {
	q := repoImpl.db
	if title != "" {
		strings.ToLower(title)
		q = q.Where("lower(title) LIKE ?", "%"+title+"%")
	}
	if genre != "" {
		strings.ToLower(genre)
		q = q.Where("lower(genre) LIKE ?", "%"+genre+"%")
	}
	if releaseDate != "" {
		strings.ToLower(releaseDate)
		q = q.Where("release_date LIKE ?", releaseDate)
	}

	var films []Film
	err := q.Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}

// FindByID returns the first film with the given ID
func (repoImpl *filmRepositoryImpl) FindByID(id uint) (*Film, error) {
	var film Film
	err := repoImpl.db.First(&film, id).Error
	if err != nil {
		return nil, err
	}
	return &film, nil
}

// Create will insert a given film in the DB
func (repoImpl *filmRepositoryImpl) Create(film *Film) error {
	err := repoImpl.db.Create(film).Error
	if err != nil {
		return err
	}
	return nil
}

// Update will update a given film in the DB, or create it if it doesn't exist
func (repoImpl *filmRepositoryImpl) Update(film *Film) error {
	err := repoImpl.db.Save(film).Error
	if err != nil {
		return err
	}
	return nil
}
