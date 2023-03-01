package controllers

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/database"
	"jwt-golang/models"
	"time"
)

var productCollection = database.OpenCollection(database.Client, "products")

func CreateProduct(c *fiber.Ctx) error {
	gofakeit.Seed(0)
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
	//categories := gofakeit.Categories()
	product.Category = "shirts"
	product.Name = gofakeit.Name()
	product.Description = gofakeit.Sentence(10)
	product.Price = gofakeit.Price(100, 1000)
	product.AvailableQuantity = gofakeit.Number(1, 100)
	product.Images = []string{gofakeit.ImageURL(100, 100)}

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

func GetAllProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// get all products
	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error getting products",
			"data":    err,
		})
	}

	// parse products
	var products []bson.M
	if err := cursor.All(ctx, &products); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error parsing products",
			"data":    err,
		})
	}
	// return success
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Products fetched successfully",
		"data":    products,
	})
}

func GetProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// get product id from params
	id := c.Params("id")

	// parse id
	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product id",
			"data":    err,
		})
	}

	// get product
	var product models.Product
	if err := productCollection.FindOne(ctx, bson.M{"_id": productId}).Decode(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error getting product",
			"data":    err,
		})
	}

	// return product
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product fetched successfully",
		"data":    product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	// TODO
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// take update data from request body
	var updateData bson.M
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid update data for model binding",
			"data":    err,
		})
	}

	//// check and retrieve userType from token
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}
	//
	//// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can create products",
			"data":    nil,
		})
	}

	// get product id from params
	id := c.Params("id")

	// parse id
	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product id",
			"data":    err,
		})
	}

	// filter product by id
	filter := bson.M{"_id": productId}

	// update product
	product, err := productCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": updateData}).DecodeBytes()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error updating product",
			"data":    err,
		})
	}
	//return success
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
		"data":    product.String(),
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	//// check and retrieve userType from token
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}
	//
	//// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can create products",
			"data":    nil,
		})
	}

	// get product id from params
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product id",
			"data":    err,
		})
	}

	// filter product by id
	filter := bson.M{"_id": id}

	// delete product
	_, err = productCollection.FindOneAndDelete(ctx, filter).DecodeBytes()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error deleting product",
			"data":    err,
		})
	}

	// return success
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully",
		"data":    nil,
	})

}
