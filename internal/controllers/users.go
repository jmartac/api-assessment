package controllers

import (
	"api-assessment/internal/auth"
	"api-assessment/internal/models"
	"api-assessment/internal/security"
	"api-assessment/internal/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type UsersController struct {
	us services.UserService
}

// NewUsersController is used to create a new Users controller.
func NewUsersController(us services.UserService) *UsersController {
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
	exists, err := uc.us.UsernameExists(user.Username)
	if err != nil {
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}
	if exists {
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

	// generate a JWT token
	tokenInfo, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		uc.handleError(w, err, http.StatusInternalServerError)
		return
	}

	uc.writeResponse(w, tokenInfo, user.ToResponse())
}

// writeResponse will try to write the given response to the client
func (uc *UsersController) writeResponse(w http.ResponseWriter, tokenInfo auth.TokenInfo, data interface{}) {
	response := struct {
		auth.TokenInfo
		Data interface{} `json:"data"`
	}{
		TokenInfo: tokenInfo,
		Data:      data,
	}

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
