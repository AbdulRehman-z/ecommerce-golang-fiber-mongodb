package controllers

import (
	"github.com/gofiber/fiber/v2"
	"jwt-golang/database"
)

var ordersCollection = database.OpenCollection(database.Client, "orders")

type ProductOrder struct {
	ProductId  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"price"`
}

func CreateOrder(c *fiber.Ctx) error {
	return c.SendString("Create order")
}
