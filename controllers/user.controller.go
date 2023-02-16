package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/models"
	"jwt-golang/utils"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func Signup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	var user models.User
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	//u.ID = .InsertedID.(primitive.ObjectID)
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
			"data":    c.JSON(existingUser),
		})
	}

	//hash password
	password, err := utils.HashPassword(user.Password)
	user.Password = password
	// insert user into database
	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to insert a user",
			"data":    err.Error(),
		})
	}

	// sign jwt with user id and email
	signedToken, err := utils.CreateToken(user.Id, user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create token",
			"data":    err.Error(),
		})
	}

	// add token to cookie session
	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	// set cookie
	c.Cookie(cookie)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User signed up successfully",
		"data":    user,
	})
}

func Signin(c *fiber.Ctx) {

}

func Logout(c *fiber.Ctx) {

}

func Profile(c *fiber.Ctx) {

}
