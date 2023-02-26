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

	// user quantity must be less than product quantity
	fmt.Println("User quantity", productToOrder.BuyQuantity, "Available quantity", product.AvailableQuantity)
	if productToOrder.BuyQuantity > product.AvailableQuantity {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Quantity must be less than product quantity",
			"data":    nil,
		})
	}

	// find user by id
	if len(user.UserCart) == 0 {
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
	} else {
		// check if product is already in cart
		for _, item := range user.UserCart {
			boolT := item.ProductId == productId
			fmt.Println(boolT)
			if item.ProductId == productId {
				// update quantity in user cart array inside user document
				filter = bson.M{"_id": userId, "userCart.productId": productId}
				update := bson.M{"$inc": bson.M{"userCart.$.buyQuantity": productToOrder.BuyQuantity}}
				if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
					return c.Status(400).JSON(fiber.Map{
						"status":  "error",
						"message": "Failed to update product quantity",
						"data":    nil,
					})
				}
			} else {
				fmt.Println("Product quantity updated")
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
			}
		}
	}

	subtractQuantity := product.AvailableQuantity - productToOrder.BuyQuantity

	fmt.Println("User cart", user.UserCart)

	product.AvailableQuantity = subtractQuantity

	// update Available quantity in product document
	filter = bson.M{"_id": productId}
	update := bson.M{"$set": bson.M{"availableQuantity": subtractQuantity}}
	if _, err := productCollection.UpdateOne(ctx, filter, update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update product quantity",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product added to cart successfully",
		"data":    productToOrder,
	})
}

func RemoveProductFromCart(c *fiber.Ctx) error {
	type Request struct {
		Count int `json:"count" bson:"count"`
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// parse req body
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"data":    err,
		})
	}

	// get product id from params and user id from context
	productId := c.Params("id")
	pId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product id",
			"data":    err,
		})
	}

	var product models.Product
	// find product by id and update availableQuantity
	filter := bson.M{"_id": pId}
	if err = productCollection.FindOneAndUpdate(ctx, filter, bson.M{"$inc": bson.M{"availableQuantity": req.Count}}).Decode(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
			"data":    err,
		})
	}

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

	// find user by id and update userCart
	var user models.User
	filter = bson.M{"_id": userId}
	if err := userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to remove product from cart",
			"data":    err,
		})
	}
	// now find product in user cart and update quantity
	for _, item := range user.UserCart {
		if item.ProductId == pId {
			// update BuyQuantity inside an element of userCart array inside user document that matches the productId
			var update bson.M
			filter = bson.M{"_id": userId, "userCart.productId": pId}
			update = bson.M{"$inc": bson.M{"userCart.$.buyQuantity": -req.Count}}

			// Check if the resulting buyQuantity is zero and remove the product if it is
			if item.BuyQuantity-req.Count == 0 {
				update = bson.M{"$pull": bson.M{"userCart": bson.M{"productId": pId}}}
			}

			if _, err := userCollection.UpdateOne(ctx, filter, update); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"status":  "error",
					"message": "Failed to update product quantity",
					"data":    nil,
				})
			}
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product removed from cart successfully",
	})
}
