package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId         string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Username       string             `json:"username,omitempty" bson:"username,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	UserType       string             `json:"userType,omitempty" bson:"userType,omitempty"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	AddressDetails []Address          `json:"addresses,omitempty" bson:"addresses,omitempty"`
	OrderStatus    []Order            `json:"orders,omitempty" bson:"orders,omitempty"`
	UserCart       []ProductsToOrder  `json:"cart,omitempty" bson:"cart,omitempty"`
}

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

type Order struct {
	Id         string            `json:"id,omitempty" bson:"id,omitempty"`
	OrderCart  []ProductsToOrder `json:"orderCart,omitempty" bson:"orderCart,omitempty"`
	TotalPrice float64           `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	Status     bool              `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  string            `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type ProductsToOrder struct {
	ProductId string  `json:"productId,omitempty" bson:"productId,omitempty"`
	Name      string  `json:"name,omitempty" bson:"name,omitempty"`
	Quantity  int     `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty" bson:"price,omitempty"`
	Rating    int     `json:"rating,omitempty" bson:"rating,omitempty"`
}

type Address struct {
	Id      primitive.ObjectID `json:"addressId,omitempty" bson:"addressId,omitempty"`
	Zip     string             `json:"zip,omitempty" bson:"zip,omitempty"`
	City    string             `json:"city,omitempty" bson:"city,omitempty"`
	State   string             `json:"state,omitempty" bson:"state,omitempty"`
	Country string             `json:"country,omitempty" bson:"country,omitempty"`
	Street  string             `json:"street,omitempty" bson:"street,omitempty"`
	Number  string             `json:"number,omitempty" bson:"number,omitempty"`
}
