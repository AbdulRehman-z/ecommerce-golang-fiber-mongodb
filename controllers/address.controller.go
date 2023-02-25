package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/models"
	"time"
)

func RemoveAddress(c *fiber.Ctx) error {
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
	emptyAddress := make([]models.Address, 0)

	// find user by id
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"addresses": emptyAddress}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    err,
		})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Address removed successfully",
		"data":    nil,
	})

}
