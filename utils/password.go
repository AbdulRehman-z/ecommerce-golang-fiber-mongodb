package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) ([]byte, []byte, error) {

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return make([]byte, 0), make([]byte, 0), err
	}

	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return hashedPassword, salt, nil
}

// ComparePassword compares the password with the hash
func VerifyPassword() {

}

func VerifyToken(token string) {
}
