package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    string             `json:"userId"`
	FirstName string             `json:"firstname" validate:"required"`
	LastName  string             `json:"lastname" validate:"required"`
	Email     string             `json:"email" validate:"required, min=6, max=10"`
	Password  string             `json:"password" validate:"required"`
	UserType  string             `json:"userType" validate:"required, eq=ADMIN|eq=USER"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Id        string
}
