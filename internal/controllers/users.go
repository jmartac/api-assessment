package controllers

import (
	"api-assessment/internal/auth"
	"api-assessment/internal/errors"
	"api-assessment/internal/models"
	"api-assessment/internal/security"
	"api-assessment/internal/services"
	"api-assessment/internal/validators"
	"encoding/json"
	"net/http"
	"strings"
)

type UsersController struct {
	us        services.UserService
	validator *validators.Validator
}

// NewUsersController is used to create a new Users controller.
func NewUsersController(us services.UserService) *UsersController {
	return &UsersController{
		us:        us,
		validator: validators.NewValidator(),
	}
}

// Create is used to create a new user
// POST /register
func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrBadRequest)
		return
	}

	// validate the user request
	err = uc.validator.Validate(userRequest)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrUserValidationFailed)
		return
	}

	// check if the user already exists
	exists, err := uc.us.UsernameExists(userRequest.Username)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrInternal)
		return
	}
	if exists {
		uc.handleError(w, nil, apiErrors.ErrUsernameAlreadyExists)
		return
	}

	// generate a password hash and flush the password
	passwordHash, err := security.GeneratePasswordHash(userRequest.Password)
	userRequest.Password = "" // just in case
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrIncorrectPassword)
		return
	}

	user := &models.User{
		Username:     userRequest.Username,
		PasswordHash: passwordHash,
	}
	err = uc.us.Create(user)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	// generate a JWT token
	tokenInfo, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	uc.writeResponse(w, tokenInfo, user.ToResponse())
}

// Login is used to log in a user
// POST /login
func (uc *UsersController) Login(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrBadRequest)
		return
	}

	// validate the user request
	err = uc.validator.Validate(userRequest)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrUserValidationFailed)
		return
	}

	// check if the user already exists
	user, err := uc.us.FindByUsername(userRequest.Username)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrNotFound)
		return
	}

	// check password
	err = security.PasswordMatches(userRequest.Password, user.PasswordHash)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrIncorrectPassword)
		return
	}

	// generate a JWT token
	tokenInfo, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		uc.handleError(w, err, apiErrors.ErrInternal)
		return
	}

	uc.writeResponse(w, tokenInfo, user.ToResponse())
}

// writeResponse will make any necessary changes to the data and write the response to the client
func (uc *UsersController) writeResponse(w http.ResponseWriter, tokenInfo auth.TokenInfo, data interface{}) {
	response := struct {
		auth.TokenInfo
		Data interface{} `json:"data"`
	}{
		TokenInfo: tokenInfo,
		Data:      data,
	}

	writeResponse(w, response)
}

// handleError will handle the given error and write the appropriate response to the client
func (uc *UsersController) handleError(w http.ResponseWriter, err error, defaultApiError apiErrors.ApiError) {
	if err == nil {
		handleError(w, err, defaultApiError)
		return
	}

	switch {
	case strings.Contains(err.Error(), "record not found"):
		handleError(w, err, apiErrors.ErrUserNotFound)
	case strings.Contains(err.Error(), "Error 1062 (23000)"): // duplicate entry
		handleError(w, err, apiErrors.ErrFilmTitleAlreadyExists)
	default:
		handleError(w, err, defaultApiError)
	}
}
