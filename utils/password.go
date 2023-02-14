package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passowrd string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passowrd), 14)
	return string(hashedPassword)
}

func VerifyPassword() {

}

func VerifyToken() {

}
