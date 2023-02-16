package controllers

import (
	"api-assessment/internal/errors"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// getIdFromPath will return the ID from the URL path
// e.g. /films/1 will return 1
func getIdFromPath(r *http.Request) (uint, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// handleError will log the error and return a generic error message to the client
func handleError(w http.ResponseWriter, err error, apiError apiErrors.ApiError) {
	log.Println(err)
	w.WriteHeader(apiError.StatusErrCode)

	writeResponse(w, apiError.ToResponse())
}

// writeResponse will try to write the given response to the client.
// If it fails it will return a generic error message
func writeResponse(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
