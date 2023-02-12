package middleware

import (
	"api-assessment/internal/auth"
	"context"
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
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// check if token is valid
		jwtClaims, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// add the claims to the context, so that other handlers can access them later
		ctxWithClaims := context.WithValue(r.Context(), auth.JWTClaimsKey, jwtClaims)
		r = r.WithContext(ctxWithClaims)

		handlerFunc.ServeHTTP(w, r)
	}
}
