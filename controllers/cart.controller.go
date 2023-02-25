package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jwt-golang/models"
	"time"
)

func AddProductToCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get user id from context
	idLocal := c.Locals("id").(string)

	userId, err := primitive.ObjectIDFromHex(idLocal)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user id",
			"data":    err,
		})
	}
	// get item id from body
	id := c.Params("id")
	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product id",
			"data":    err,
		})
	}
	// find product by id
	filter := bson.M{"_id": productId}

	//var product models.ProductsToOrder
	var product models.Product
	err = productCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
			"data":    err,
		})
	}

	// check if product is available
	if product.AvailableQuantity == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Product is not available",
			"data":    nil,
		})
	}

	// fetch UserCart from user
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    err,
		})
	}

	// parse req body
	var productToOrder models.ProductsToOrder
	productToOrder.ProductId = productId
	productToOrder.CreatedAt = time.Now()
	productToOrder.UpdatedAt = time.Now()
	if err := c.BodyParser(&productToOrder); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"data":    err,
		})
	}

	//// convert quantity to int
	////quantityInt, err = strconv.Atoi(string(quantityInt))
	//if err != nil {
	//	return c.Status(400).JSON(fiber.Map{
	//		"status":  "error",
	//		"message": "Invalid quantity",
	//		"data":    err,
	//	})
	//}

	//// fetch the product from user Product collection
	//var product models.Product
	//if err = productCollection.FindOne(ctx, bson.M{"_id": productId}).Decode(&product); err != nil {
	//	return c.Status(400).JSON(fiber.Map{
	//		"status":  "error",
	//		"message": "Product not found",
	//		"data":    err,
	//	})
	//}

	// user quantity must be less than product quantity
	fmt.Println("User quantity", productToOrder.BuyQuantity, "Available quantity", product.AvailableQuantity)
	if productToOrder.BuyQuantity > product.AvailableQuantity {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Quantity must be less than product quantity",
			"data":    nil,
		})
	}

	// exclude some fields from fetched product
	//userCartProduct := &UserCartProduct{
	//	ID:   productToOrder.ProductId,
	//	Name: productToOrder.Name,
	//	Price:    productToOrder.Price,
	//	Quantity: quantityInt,
	//}

	// find user by id
	filter = bson.M{"_id": userId}
	update := bson.M{"$push": bson.M{"userCart": &productToOrder}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to add product to cart",
			"data":    err,
		})
	}

	subtractQuantity := product.AvailableQuantity - productToOrder.BuyQuantity

	fmt.Println("User cart", user.UserCart)

	product.AvailableQuantity = subtractQuantity

	// update Available
	filter = bson.M{"_id": productId}
	update = bson.M{"$set": bson.M{"availableQuantity": subtractQuantity}}
	if _, err := productCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update product quantity",
			"data":    nil,
		})
	}

	//
	//for _, cartProduct := range user.UserCart {
	//	if cartProduct.ProductId == productId {
	//		// update AvailableQuantity field for original product in product collection
	//		filter = bson.M{"_id": productId}
	//		update := bson.M{"$set": bson.M{"availableQuantity": subtractQuantity}}
	//		if _, err := productCollection.UpdateOne(ctx, filter, update); err != nil {
	//			return c.Status(400).JSON(fiber.Map{
	//				"status":  "error",
	//				"message": "Product already in cart",
	//				"data":    nil,
	//			})
	//		}
	//	}
	//}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product added to cart successfully",
		"data":    productToOrder,
	})
}
