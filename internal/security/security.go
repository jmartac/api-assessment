package security

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

var pepper string

func init() {
	pepper = os.Getenv("HASH_PEPPER")
}

// GeneratePasswordHash generates a password hash from a peppered password
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// PasswordMatches checks whether the plaintext password hash equals to given hashedPassword or not.
// Returns nil on success, otherwise returns an error
func PasswordMatches(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+pepper))
}
