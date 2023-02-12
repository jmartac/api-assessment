package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

const expirationTime = 30 * time.Minute
const JWTClaimsKey = "jwtClaims"

var jwtSecret = []byte("df83hjfr8sj39gtyuhw93fc598mn7") // TODO move to env variables

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
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// keyFunc checks the signing method and returns the secret key
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		err := errors.New("signing method family different from HMAC-SHA")
		return nil, err
	}
	return jwtSecret, nil
}
