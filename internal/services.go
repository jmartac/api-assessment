package internal

import (
	"api-assessment/internal/models"
	"api-assessment/internal/services"
	"gorm.io/gorm"
)

// NewServices creates and returns a new Services struct
func NewServices() (*Services, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}

	return &Services{
		db:          db,
		FilmService: services.NewFilmService(db),
		UserService: services.NewUserService(db),
	}, nil
}

// Services contains all the services used by the application, including the database connection
type Services struct {
	db          *gorm.DB
	FilmService services.FilmService
	UserService services.UserService
}

// Close closes the database connection
func (s *Services) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// AutoMigrate will attempt to migrate all models tables
func (s *Services) AutoMigrate() error {
	return s.db.Migrator().AutoMigrate(&models.Film{}, &models.User{})
}
