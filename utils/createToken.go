package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

type Claims struct {
	Id    primitive.ObjectID
	Email string
	jwt.StandardClaims
}

func CreateToken(id primitive.ObjectID, email string) (tokenString string, err error) {
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
