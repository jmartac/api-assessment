package models

import (
	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Title       string `gorm:"not_null;unique"`
	Director    string
	ReleaseDate string
	Genre       string
	Synopsis    string
	Cast        string
	UserID      uint
	User        User // with this, gorm is able to understand the dependency between films and users
}

// MergeData merges non-empty fields from the given data into the film
func (f *Film) MergeData(data *FilmRequest) {
	if data.Title != "" {
		f.Title = data.Title
	}
	if data.Director != "" {
		f.Director = data.Director
	}
	if data.ReleaseDate != "" {
		f.ReleaseDate = data.ReleaseDate
	}
	if data.Genre != "" {
		f.Genre = data.Genre
	}
	if data.Synopsis != "" {
		f.Synopsis = data.Synopsis
	}
	if data.Cast != "" {
		f.Cast = data.Cast
	}
}

// ToResponse converts a Film model into a FilmResponse
func (f *Film) ToResponse() FilmResponse {
	return FilmResponse{
		ID:          f.ID,
		Title:       f.Title,
		Director:    f.Director,
		ReleaseDate: f.ReleaseDate,
		Genre:       f.Genre,
		Synopsis:    f.Synopsis,
		Cast:        f.Cast,
		UserID:      f.UserID,
	}
}

type Films []Film

// ToResponse converts a slice of Film into a slice of FilmResponse
func (f Films) ToResponse() []FilmResponse {
	films := make([]FilmResponse, 0)
	for _, film := range f {
		films = append(films, film.ToResponse())
	}
	return films
}

type FilmRequest struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Synopsis    string `json:"synopsis"`
	Cast        string `json:"cast"`
}

type FilmResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Synopsis    string `json:"synopsis"`
	Cast        string `json:"cast"`
	UserID      uint   `json:"user_id"`
}
