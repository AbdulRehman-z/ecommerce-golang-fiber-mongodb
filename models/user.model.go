package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    string             `json:"userId,omitempty" bson:"userId,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Hash      []byte             `json:"hash,omitempty" bson:"hash,omitempty"`
	Salt      []byte             `json:"salt,omitempty" bson:"salt,omitempty"`
	UserType  string             `json:"userType,omitempty" bson:"userType,omitempty"`
	Id        string             `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
