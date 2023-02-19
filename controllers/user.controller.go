package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/models"
	"jwt-golang/utils"
	"os"
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
	user.UserId = user.ID.Hex()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    err.Error(),
		})
	}

	// check if individual is USER or ADMIN
	if user.Password == os.Getenv("ADMIN_PASS") && user.Email == os.Getenv("ADMIN_EMAIL") {
		user.UserType = "ADMIN"
	} else {
		user.UserType = "USER"
	}
	fmt.Println(user.UserType)
	// allow only one admin user
	if user.UserType == "ADMIN" {
		filter := bson.M{"userType": "ADMIN"}
		if _, err := userCollection.FindOne(ctx, filter).DecodeBytes(); err == nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "Admin user already exists",
				"data":    c.JSON(err),
			})
		}
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
	signedToken, err := utils.CreateToken(user.ID, user.Email)
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

func Signin(c *fiber.Ctx) error {

	type SigninRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// parse request body
	var req SigninRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"data":    err.Error(),
		})
	}

	// get email and password from request body
	password := req.Password
	email := req.Email

	// check if user exists
	existingUser, err := userCollection.FindOne(ctx, bson.M{"email": email}).DecodeBytes()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User does not exist",
			"data":    err.Error(),
		})
	}

	// check if password is correct
	isValid := utils.VerifyPassword(password, existingUser.Lookup("password").StringValue())
	if !isValid {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials",
			"data":    nil,
		})
	}

	// sign jwt with user id and email
	signedToken, err := utils.CreateToken(existingUser.Lookup("_id").ObjectID(), existingUser.Lookup("email").StringValue())

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
		"message": "User signed in successfully",
		"data":    existingUser,
	})
}

func Signout(c *fiber.Ctx) error {
	// delete cookie
	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	// set cookie
	c.Cookie(cookie)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged out successfully",
		"data":    nil,
	})
}

func Profile(c *fiber.Ctx) error {
	// get id and email from fiber.context
	// this is more efficient
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to get id",
		})
	}

	// get user from database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingUser, err := userCollection.FindOne(ctx, bson.M{"_id": id}).DecodeBytes()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "user does not exist",
			"data":    err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully fetched user",
		"data":    existingUser,
	})
}
