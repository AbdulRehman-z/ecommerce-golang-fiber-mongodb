package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/models"
	"time"
)

func UpdateAddress(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idLocal := c.Locals("id").(string)
	userId, err := primitive.ObjectIDFromHex(idLocal)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user id",
			"data":    err,
		})
	}
	// get address from body
	var address models.Address
	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid address",
			"data":    err,
		})
	}

	// find user by id and update address
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"address": address}}
	if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Address not updated",
			"data":    err,
		})
	}

	// return success message
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Address updated successfully",
		"data":    nil,
	})

}
