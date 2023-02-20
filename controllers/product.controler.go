package controllers

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/models"
	"time"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func CreateProduct(c *fiber.Ctx) error {

	gofakeit.Seed(11)
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
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
	}

	// create product model
	var product models.Product
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	//	parse request body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product data for model binding",
			"data":    err,
		})
	}

	// check if product already exists
	filter := bson.M{"name": product.ID}
	if existingProduct, err := productCollection.FindOne(ctx, filter).DecodeBytes(); err == nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Product already exists",
			"data":    existingProduct,
		})
	}

	// insert product
	if _, err := productCollection.InsertOne(ctx, product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error inserting product",
			"data":    err,
		})
	}

	// return success
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
		"data":    product,
	})
}

func GetAllProducts(c *fiber.Ctx) {

}

func GetProduct(c *fiber.Ctx) {

}

func UpdateProduct(c *fiber.Ctx) {

}

func DeleteProduct(c *fiber.Ctx) {

}
