package apiErrors

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound     = ApiError{errors.New("resource not found"), "E0000", http.StatusNotFound}
	ErrInvalidID    = ApiError{errors.New("invalid id"), "E0001", http.StatusBadRequest}
	ErrBadRequest   = ApiError{errors.New("the request is invalid"), "E0002", http.StatusBadRequest}
	ErrUnauthorized = ApiError{errors.New("user is not authorized to perform this action"), "E0003", http.StatusUnauthorized}
	ErrInternal     = ApiError{errors.New("internal server error"), "E0004", http.StatusInternalServerError}

	// Users

	ErrIncorrectPassword     = ApiError{errors.New("incorrect password"), "E0010", http.StatusBadRequest}
	ErrUsernameAlreadyExists = ApiError{errors.New("username already exists"), "E0011", http.StatusBadRequest}
	ErrUserValidationFailed  = ApiError{errors.New("validation errors: " +
		"the username (3 to 30 chars) must start with a letter, it can only contain letters and numbers; " +
		"the password (8 to 128 chars) must start with a letter, it can only contain letters and numbers"), "E0012", http.StatusBadRequest}
	ErrUserNotFound = ApiError{errors.New("user not found"), "E0013", http.StatusNotFound}

	// Films

	ErrFilmTitleAlreadyExists = ApiError{errors.New("film title already exists"), "E0020", http.StatusBadRequest}
	ErrFilmNotFound           = ApiError{errors.New("film not found"), "E0021", http.StatusNotFound}
	ErrFilmTitleRequired      = ApiError{errors.New("film title is required"), "E0022", http.StatusBadRequest}

	// Auth

	ErrInvalidToken = ApiError{errors.New("missing or invalid token"), "E0030", http.StatusUnauthorized}
	ErrAuthFailed   = ApiError{errors.New("authentication is required to access this resource"), "E0031", http.StatusUnauthorized}
	ErrTokenExpired = ApiError{errors.New("token is expired"), "E0032", http.StatusUnauthorized}
)

type ApiError struct {
	error
	ErrCode       string
	StatusErrCode int
}

func (e ApiError) Error() string {
	return e.error.Error()
}

func (e ApiError) ToResponse() ErrorResponse {
	return ErrorResponse{
		ErrCode: e.ErrCode,
		ErrMsg:  fmt.Sprintf("%s: %s", http.StatusText(e.StatusErrCode), e.Error()),
	}
}

type ErrorResponse struct {
	ErrCode string `json:"code"`
	ErrMsg  string `json:"error"`
}
