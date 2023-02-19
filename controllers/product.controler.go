package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func CreateProduct(c *fiber.Ctx) {
	// check if user is signed in and user is admin
	//var name string
	// if user is not admin, return error
	// if user is not admin, return error

}

func GetAllProducts(c *fiber.Ctx) {

}

func GetProduct(c *fiber.Ctx) {

}

func UpdateProduct(c *fiber.Ctx) {

}

func DeleteProduct(c *fiber.Ctx) {

}
