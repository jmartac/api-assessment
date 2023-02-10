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
	FindByID(id uint) (*Film, error)
	Create(film *Film) error
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
func (repoImpl *filmRepositoryImpl) FindAll() ([]Film, error) {
	var films []Film
	err := repoImpl.db.Find(&films).Error
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
