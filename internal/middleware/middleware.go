package middleware

import (
	"api-assessment/internal/auth"
	"api-assessment/internal/errors"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// LogAPICalls is a middleware to log all API calls
func LogAPICalls(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

// JsonHeader is a middleware to set the response Content-Type header to application/json
func JsonHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

// Authenticate is a middleware to authenticate a request via JWT
func Authenticate(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check auth header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			handleError(w, apiErrors.ErrInvalidToken)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// check if token is valid
		jwtClaim, err := auth.ValidateToken(token)
		if err != nil {
			var apiError apiErrors.ApiError
			if strings.Contains(err.Error(), "token is expired by") {
				apiError = apiErrors.ErrTokenExpired
			} else {
				apiError = apiErrors.ErrInvalidToken
			}

			handleError(w, apiError)
			return
		}

		// add the claims to the context, so that other handlers can access them later
		ctxWithClaims := context.WithValue(r.Context(), auth.JWTClaimsKey, jwtClaim)
		r = r.WithContext(ctxWithClaims)

		handlerFunc.ServeHTTP(w, r)
	}
}

func handleError(w http.ResponseWriter, apiError apiErrors.ApiError) {
	w.WriteHeader(apiError.StatusErrCode)
	err := json.NewEncoder(w).Encode(apiError.ToResponse())

	if err != nil {
		log.Println(err)
		http.Error(w, apiErrors.ErrInternal.Error(), apiErrors.ErrInternal.StatusErrCode)
	}
}
