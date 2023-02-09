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
	log.Println("GET /films")
	w.Header().Set("Content-Type", "application/json")

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
	log.Println("GET /films/{id}")
	w.Header().Set("Content-Type", "application/json")

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

func (fc *FilmsController) extractID(r *http.Request) (uint, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
