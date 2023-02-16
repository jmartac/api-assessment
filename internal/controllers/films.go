package controllers

import (
	"api-assessment/internal/auth"
	"api-assessment/internal/errors"
	"api-assessment/internal/models"
	"api-assessment/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

type FilmsController struct {
	fs services.FilmService
}

// NewFilmsController is used to create a new Films controller.
func NewFilmsController(fs services.FilmService) *FilmsController {
	return &FilmsController{fs: fs}
}

// FindAll is used to find all films
// GET /films
func (fc *FilmsController) FindAll(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	genre := r.URL.Query().Get("genre")
	releaseDate := r.URL.Query().Get("release_date")

	films, err := fc.fs.FindAll(title, genre, releaseDate)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	fc.writeResponse(w, models.Films(films).ToResponse())
}

// FindByID is used to find the details of a film by ID
// GET /films/{id}
func (fc *FilmsController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromPath(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInvalidID)
		return
	}

	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Create is used to create a new film and return the details
// POST /films
func (fc *FilmsController) Create(w http.ResponseWriter, r *http.Request) {
	var filmRequest models.FilmRequest
	err := json.NewDecoder(r.Body).Decode(&filmRequest)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrBadRequest)
		return
	}

	// validate that film title is not empty
	if strings.TrimSpace(filmRequest.Title) == "" {
		fc.handleError(w, nil, apiErrors.ErrFilmTitleRequired)
		return
	}

	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrAuthFailed)
		return
	}

	film := models.Film{
		Title:       filmRequest.Title,
		Director:    filmRequest.Director,
		ReleaseDate: filmRequest.ReleaseDate,
		Genre:       filmRequest.Genre,
		Synopsis:    filmRequest.Synopsis,
		Cast:        filmRequest.Cast,
		UserID:      userID,
	}
	err = fc.fs.Create(&film)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Update is used to update a given film and return the updated details
// POST /films/{id}/update
func (fc *FilmsController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromPath(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInvalidID)
		return
	}

	// Retrieve film from database
	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	// check if user is authorized to update film
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrAuthFailed)
		return
	}
	if userID != film.UserID {
		fc.handleError(w, err, apiErrors.ErrUnauthorized)
		return
	}

	// Decode new data
	var filmRequest models.FilmRequest
	err = json.NewDecoder(r.Body).Decode(&filmRequest)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrBadRequest)
		return
	}

	// merge new data with existing film data
	film.MergeData(&filmRequest)

	err = fc.fs.Update(film)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Delete is used to delete the film with the given ID
// POST /films/{id}/delete
func (fc *FilmsController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromPath(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInvalidID)
		return
	}

	// retrieve film from database
	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	// check if user is authorized to delete film
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrAuthFailed)
		return
	}
	if userID != film.UserID {
		fc.handleError(w, err, apiErrors.ErrUnauthorized)
		return
	}

	err = fc.fs.Delete(id)
	if err != nil {
		fc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{"Film deleted successfully"}

	fc.writeResponse(w, response)
}

// writeResponse will make any necessary changes to the data and write the response to the client
func (fc *FilmsController) writeResponse(w http.ResponseWriter, data interface{}) {
	// this controller doesn't need to make any changes to the data
	writeResponse(w, data)
}

// handleError will handle the given error and write the appropriate response to the client
func (fc *FilmsController) handleError(w http.ResponseWriter, err error, defaultApiError apiErrors.ApiError) {
	if err == nil {
		handleError(w, err, defaultApiError)
		return
	}

	switch {
	case strings.Contains(err.Error(), "record not found"):
		handleError(w, err, apiErrors.ErrFilmNotFound)
	case strings.Contains(err.Error(), "Error 1062 (23000)"): // duplicate entry
		handleError(w, err, apiErrors.ErrFilmTitleAlreadyExists)
	default:
		handleError(w, err, defaultApiError)
	}
}
