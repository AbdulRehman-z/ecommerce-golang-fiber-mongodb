package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/models"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func Signup(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data",
			"data":    err.Error(),
		})
	}

	// check if user already exists
	filter := bson.M{"email": user.Email}
	if existingUser, err := userCollection.FindOne(ctx, filter).DecodeBytes(); err == nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
			"data":    existingUser,
		})
	}

	// insert user into database
	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to insert a user",
			"data":    err.Error(),
		})
	}

	// sign jwt with user id and email

}

func Signin(c *fiber.Ctx) error {

}

func Logout(c *fiber.Ctx) error {

}

func Profile(c *fiber.Ctx) error {

}
