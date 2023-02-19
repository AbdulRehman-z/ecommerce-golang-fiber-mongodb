package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"time"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func CreateProduct(c *fiber.Ctx) error {

	_, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	// check if user is signed in
	userType, err := c.Locals("userType").(string)
	if !err {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    err,
		})
	}

	// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can create products",
			"data":    err,
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "You are admin",
			"data":    nil,
		})
	}

}

// if user is admin, create product

func GetAllProducts(c *fiber.Ctx) {

}

func GetProduct(c *fiber.Ctx) {

}

func UpdateProduct(c *fiber.Ctx) {

}

func DeleteProduct(c *fiber.Ctx) {

}
