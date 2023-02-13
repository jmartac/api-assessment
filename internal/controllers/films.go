package controllers

import (
	"api-assessment/internal/auth"
	"api-assessment/internal/models"
	"api-assessment/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, models.Films(films).ToResponse())
}

// FindByID is used to find the details of a film by ID
// GET /films/{id}
func (fc *FilmsController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := fc.extractID(r)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Create is used to create a new film and return the details
// POST /films
func (fc *FilmsController) Create(w http.ResponseWriter, r *http.Request) {
	var request models.FilmRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		fc.handleError(w, err, http.StatusUnauthorized)
		return
	}

	film := models.Film{
		Title:       request.Title,
		Director:    request.Director,
		ReleaseDate: request.ReleaseDate,
		Genre:       request.Genre,
		Synopsis:    request.Synopsis,
		UserID:      userID,
	}
	err = fc.fs.Create(&film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Update is used to update a given film and return the updated details
// POST /films/{id}/update
func (fc *FilmsController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := fc.extractID(r)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	// Retrieve film from database
	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, http.StatusNotFound)
		return
	}

	// check if user is authorized to update film
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil || userID != film.UserID {
		fc.handleError(w, err, http.StatusUnauthorized)
		return
	}

	// Decode new data
	var request models.FilmRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	// merge new data with existing film data
	film.MergeData(&request)

	err = fc.fs.Update(film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, film.ToResponse())
}

// Delete is used to delete the film with the given ID
// POST /films/{id}/delete
func (fc *FilmsController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := fc.extractID(r)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	// retrieve film from database
	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, http.StatusNotFound)
		return
	}

	// check if user is authorized to delete film
	userID, err := auth.GetUserIDFromRequest(r)
	if err != nil || userID != film.UserID {
		fc.handleError(w, err, http.StatusUnauthorized)
		return
	}

	err = fc.fs.Delete(id)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, "Film deleted")
}

// extractID will return the ID from the URL path
func (fc *FilmsController) extractID(r *http.Request) (uint, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// writeResponse will try to write the given response to the client
func (fc *FilmsController) writeResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fc.handleError(w, err, http.StatusInternalServerError)
	}
}

// handleError will log the error and write the given status code to the client
func (fc *FilmsController) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
