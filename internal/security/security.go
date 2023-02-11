package security

import "golang.org/x/crypto/bcrypt"

const pepper = "secret-random.string" // TODO move to env variables

// GeneratePasswordHash generates a password hash from a peppered password
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
