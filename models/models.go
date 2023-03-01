package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	UserType  string             `json:"userType" bson:"userType"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Address   Address            `json:"address" bson:"address"`
	Orders    []Order            `json:"orders" bson:"orders"`
	UserCart  []ProductsToOrder  `json:"userCart" bson:"userCart"`
}

type Product struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name,omitempty" bson:"name,omitempty"`
	Price             float64            `json:"price,omitempty" bson:"price,omitempty"`
	Description       string             `json:"description,omitempty" bson:"description,omitempty"`
	AvailableQuantity int                `json:"availableQuantity" bson:"availableQuantity"`
	Category          string             `json:"category,omitempty" bson:"category,omitempty"`
	Images            []string           `json:"images,omitempty" bson:"images,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Order struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderCart  []ProductsToOrder  `json:"orderCart,omitempty" bson:"orderCart,omitempty"`
	TotalPrice float64            `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type ProductsToOrder struct {
	ProductId   primitive.ObjectID `json:"productId,omitempty" bson:"productId,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	BuyQuantity int                `json:"buyQuantity" bson:"buyQuantity"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Address struct {
	ZipCode     string `json:"zipCode,omitempty" bson:"zipCode,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	State       string `json:"state,omitempty" bson:"state,omitempty"`
	Country     string `json:"country,omitempty" bson:"country,omitempty"`
	Street      string `json:"street,omitempty" bson:"street,omitempty"`
	HouseNumber string `json:"houseNumber,omitempty" bson:"houseNumber,omitempty"`
}
