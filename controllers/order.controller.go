package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/models"
	"time"
)

func OrderOne(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.Locals("id").(string)
	uid, _ := primitive.ObjectIDFromHex(userId)
	var model models.User
	filter := bson.M{"_id": uid}
	if err := userCollection.FindOne(ctx, filter).Decode(&model); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "order does not exist",
		})
	}

	// get order id from params
	id := c.Params("id")
	orderId, _ := primitive.ObjectIDFromHex(id)

	// find order by id in userCart array
	var productToOrder models.ProductsToOrder
	for _, item := range model.UserCart {
		if item.ProductId == orderId {
			productToOrder = item
			filter := bson.M{"_id": uid, "userCart.productId": orderId}
			update := bson.M{"$pull": bson.M{"userCart": bson.M{"productId": orderId}}}
			if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"status":  "error",
					"message": "order does not exist",
				})
			}
		}

	}

	// add productToOrder to orders array
	var order models.Order
	order.Id = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.TotalPrice = productToOrder.Price * float64(productToOrder.BuyQuantity)
	order.OrderCart = append(order.OrderCart, productToOrder)

	filter = bson.M{"_id": uid}
	update := bson.M{"$push": bson.M{"orders": order}}
	if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "order does not exist",
		})
	}

	// return success message
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "order added successfully",
		"data":    productToOrder,
	})

}

func OrderAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.Locals("id").(string)
	uid, _ := primitive.ObjectIDFromHex(userId)
	var user models.User
	filter := bson.M{"_id": uid}
	if err := userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "order does not exist",
		})
	}

	// get all products from userCart array & add them to the orders array
	var orders []models.Order
	for _, item := range user.UserCart {
		var order models.Order
		order.Id = primitive.NewObjectID()
		order.CreatedAt = time.Now()
		order.TotalPrice = item.Price * float64(item.BuyQuantity)
		order.OrderCart = append(order.OrderCart, item)
		orders = append(orders, order)
	}

	// delete all products from userCart array
	filter = bson.M{"_id": uid}
	update := bson.M{"$set": bson.M{"userCart": []models.ProductsToOrder{}}}
	if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to set userCart to empty array",
		})
	}

	// add orders to orders array
	filter = bson.M{"_id": uid}
	update = bson.M{"$push": bson.M{"orders": bson.M{"$each": orders}}}
	if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to add orders to orders array",
		})
	}

	// return success message
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "orders added successfully",
		"data":    orders,
	})

}

