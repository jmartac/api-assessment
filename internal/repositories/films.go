package repositories

import (
	"api-assessment/internal/models"
	"gorm.io/gorm"
)

type FilmRepository interface {
	FindAll(title, genre, releaseDate string) ([]models.Film, error)
	FindByID(id uint) (*models.Film, error)
	Create(film *models.Film) error
	Update(film *models.Film) error
	Delete(id uint) error
}

// NewFilmRepository creates and returns a new FilmRepository
func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepositoryImpl{
		db: db,
	}
}

type filmRepositoryImpl struct {
	db *gorm.DB
}

// FindAll returns all films in the database
func (repoImpl *filmRepositoryImpl) FindAll(title, genre, releaseDate string) ([]models.Film, error) {
	q := repoImpl.db
	if title != "" {
		q = q.Where("title LIKE ?", "%"+title+"%")
	}
	if genre != "" {
		q = q.Where("genre LIKE ?", "%"+genre+"%")
	}
	if releaseDate != "" {
		q = q.Where("release_date LIKE ?", releaseDate)
	}

	var films []models.Film
	err := q.Find(&films).Error
	if err != nil {
		return nil, err
	}
	return films, nil
}

// FindByID returns the first film with the given ID
func (repoImpl *filmRepositoryImpl) FindByID(id uint) (*models.Film, error) {
	var film models.Film
	err := repoImpl.db.First(&film, id).Error
	if err != nil {
		return nil, err
	}
	return &film, nil
}

// Create will insert a given film in the DB
func (repoImpl *filmRepositoryImpl) Create(film *models.Film) error {
	err := repoImpl.db.Create(film).Error
	if err != nil {
		return err
	}
	return nil
}

// Update will update a given film in the DB, or create it if it doesn't exist
func (repoImpl *filmRepositoryImpl) Update(film *models.Film) error {
	err := repoImpl.db.Save(film).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete will delete the film with the given ID (
func (repoImpl *filmRepositoryImpl) Delete(id uint) error {
	film := &models.Film{Model: gorm.Model{ID: id}}
	err := repoImpl.db.Delete(film).Error
	if err != nil {
		return err
	}
	return nil
}
