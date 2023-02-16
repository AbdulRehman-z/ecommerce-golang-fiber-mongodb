package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword compares the password with the hash
func VerifyPassword() {

}

func VerifyToken(token string) {
}
