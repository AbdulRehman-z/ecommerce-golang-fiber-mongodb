package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/database"
	"jwt-golang/models"
	"time"
)

var collection = database.OpenCollection(database.Client, "users")

func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// check and retrieve user type from context
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}

	// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can get users",
			"data":    nil,
		})
	}

	// get id from url params
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user id",
			"data":    err,
		})
	}

	// find user by id
	var user models.User
	err = collection.FindOne(ctx, models.User{ID: id}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    err,
		})
	}

	// return user
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// check and retrieve user type from context
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}

	// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can get users",
			"data":    nil,
		})
	}

	// find all users
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error getting users",
			"data":    err,
		})
	}

	// create slice of users
	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error creating users slice",
			"data":    err,
		})
	}

	// return users
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Users found",
		"data":    users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// check and retrieve user type from context
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}

	// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can get users",
			"data":    nil,
		})
	}

	// get id from url params
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user id",
			"data":    err,
		})
	}

	// filter user
	filter := bson.M{"_id": id}

	// get user
	var user models.User
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    err,
		})
	}

	// delete user
	if _, err := userCollection.FindOneAndDelete(ctx, filter).DecodeBytes(); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error deleting user",
			"data":    err,
		})
	}

	// return success
	if user.UserType == "ADMIN" {
		// delete cookie
		cookie := &fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}
		// set empty cookie
		c.Cookie(cookie)

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Admin deleted",
			"data":    nil,
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "User deleted",
			"data":    nil,
		})
	}
}

func DeleteAllUsers(c *fiber.Ctx) error {
	// delete all users except admin
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// check and retrieve user type from context
	userType := c.Locals("userType").(string)
	if userType == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user data for model binding",
			"data":    nil,
		})
	}

	// check if user is admin
	if userType != "ADMIN" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Only admin can get users",
			"data":    nil,
		})
	}

	// delete all users except admin
	if _, err := userCollection.DeleteMany(ctx, bson.M{"userType": bson.M{"$ne": "ADMIN"}}); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error deleting users",
			"data":    err,
		})
	}

	// return success
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Users deleted",
		"data":    nil,
	})
}