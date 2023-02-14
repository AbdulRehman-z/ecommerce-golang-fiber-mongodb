package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/helpers"
	"jwt-golang/models"
	"time"
)

var collection *mongo.Collection = database.OpenCollection(database.Client, "users")

func GetAllUsers(c *fiber.Ctx) error {

}

func DeleteUser(c *fiber.Ctx) error {

}

func DeleteAllUsers(c *fiber.Ctx) error {

}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, models.User{Id: userId}).Decode(&user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Status(200).JSON(fiber.Map{"user": user})
}
