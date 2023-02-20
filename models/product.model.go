package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
