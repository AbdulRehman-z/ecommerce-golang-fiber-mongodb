package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {

	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	// hash password with argon2
	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)

	//  encode salt and hashed password to base64
	encodedPassword := fmt.Sprintf("%s.%s", base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hashedPassword))
	return encodedPassword, nil
}

// ComparePassword compares the password with the hash
func VerifyPassword() {

}

func VerifyToken(token string) {
}
