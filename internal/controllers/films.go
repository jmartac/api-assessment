package controllers

import (
	"api-assessment/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type FilmsController struct {
	fs models.FilmService
}

// NewFilmsController is used to create a new Films controller.
func NewFilmsController(fs models.FilmService) *FilmsController {
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

	fc.writeResponse(w, films)
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

	fc.writeResponse(w, film)
}

// Create is used to create a new film and return the details
// POST /films
func (fc *FilmsController) Create(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = fc.fs.Create(&film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, film)
}

// Update is used to update a given film and return the updated details
// POST /films/{id}/update
func (fc *FilmsController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := fc.extractID(r)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	film, err := fc.fs.FindByID(id)
	if err != nil {
		fc.handleError(w, err, http.StatusNotFound)
		return
	}

	// Decode new data into retrieved film
	err = json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	err = fc.fs.Update(film)
	if err != nil {
		fc.handleError(w, err, http.StatusBadRequest)
		return
	}

	fc.writeResponse(w, film)
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
func (fc *FilmsController) writeResponse(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fc.handleError(w, err, http.StatusInternalServerError)
	}
}

// handleError will log the error and write the given status code to the client
func (fc *FilmsController) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
