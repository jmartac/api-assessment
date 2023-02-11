package controllers

import (
	"api-assessment/internal/models"
	"api-assessment/internal/security"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type UsersController struct {
	us models.UserService
}

// NewUsersController is used to create a new Users controller.
func NewUsersController(us models.UserService) *UsersController {
	return &UsersController{us: us}
}

// Create is used to create a new user
// POST /register
func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		uc.handleError(w, err, http.StatusBadRequest)
		return
	}

	// check if the user already exists
	ok, err := uc.us.UsernameExists(user.Username)
	if err != nil {
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}
	if ok {
		uc.handleError(w, errors.New("user tried to register with an already exising username"), http.StatusBadRequest)
		return
	}

	// generate a password hash and flush the password
	user.PasswordHash, err = security.GeneratePasswordHash(user.Password)
	user.Password = ""
	if err != nil {
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}

	err = uc.us.Create(&user)
	if err != nil {
		uc.handleError(w, err, http.StatusBadRequest)
		return
	}

	// TODO should return a JWT

	uc.writeResponse(w, user)
}

// writeResponse will try to write the given response to the client
func (uc *UsersController) writeResponse(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		uc.handleError(w, err, http.StatusInternalServerError)
	}
}

// handleError will log the error and write the given status code to the client
func (uc *UsersController) handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
