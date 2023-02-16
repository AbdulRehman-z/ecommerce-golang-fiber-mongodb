package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type Claims struct {
	Id    string
	Email string
	jwt.StandardClaims
}

func CreateToken(id string, email string) (tokenString string, err error) {
	claims := &Claims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
