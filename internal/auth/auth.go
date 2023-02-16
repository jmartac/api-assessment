package auth

import (
	"api-assessment/internal/errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	JWTClaimsKey   = claimsKey(iota)
	expirationTime = 30 * time.Minute
)

var jwtSecret []byte

type claimsKey int

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

// TokenInfo represents a token and its expiration time
type TokenInfo struct {
	AccessToken string    `json:"access_token"`
	Expiration  time.Time `json:"expiresAt"`
}

// JWTClaim represents the claims of a JWT token
type JWTClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// GenerateToken generates a JWT token for the given user
func GenerateToken(userId uint, username string) (TokenInfo, error) {
	// token expiration time
	exp := time.Now().Add(expirationTime)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   strconv.Itoa(int(userId)),
		},
		Username: username,
	}

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{
		AccessToken: signedToken,
		Expiration:  exp,
	}, nil
}

// ValidateToken validates the given token and returns the claims if the token is valid
func ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, keyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return token.Claims.(*JWTClaim), nil
	}
	return nil, apiErrors.ErrInvalidToken
}

// GetUserIDFromRequest returns the ID of the user from the request context
func GetUserIDFromRequest(r *http.Request) (uint, error) {
	ctxValue := r.Context().Value(JWTClaimsKey)
	jwtClaim, ok := ctxValue.(*JWTClaim)
	if !ok {
		return 0, apiErrors.ErrAuthFailed
	}

	id, err := strconv.ParseUint(jwtClaim.Subject, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// keyFunc checks the signing method and returns the secret key
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, apiErrors.ErrInvalidToken
	}
	return jwtSecret, nil
}
