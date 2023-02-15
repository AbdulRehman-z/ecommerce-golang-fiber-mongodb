package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passowrd string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passowrd), 14)
	return string(hashedPassword)
}

// ComparePassword compares the password with the hash
func VerifyPassword() {

}

func VerifyToken(token string) error {
}
