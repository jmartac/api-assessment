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
func (fc *FilmsController) FindAll(w http.ResponseWriter, _ *http.Request) {
	films, err := fc.fs.FindAll()
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(films)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// FindByID is used to find the details of a film by ID
// GET /films/{id}
func (fc *FilmsController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := fc.extractID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	film, err := fc.fs.FindByID(id)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// Create is used to create a new film and return the details
// POST /films
func (fc *FilmsController) Create(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = fc.fs.Create(&film)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// extractID will return the ID from the URL path
func (fc *FilmsController) extractID(r *http.Request) (uint, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}