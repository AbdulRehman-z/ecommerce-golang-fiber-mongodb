package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"jwt-golang/utils"
)

func RequireAuthMiddleware(c *fiber.Ctx) error {
	// get token from header
	authHeader := c.Get("Authorization")
	token := c.Cookies("jwt")

	// check if token is empty
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing Authorization header",
		})
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing token",
		})
	}

	// check if token is valid
	id, email, err := utils.VerifyToken(token)
	fmt.Println("---------------------id------------------: ", id)
	fmt.Println("---------------------email-------------------: ", email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token",
		})
	}

	// set user id and email to context
	c.Locals("id", id)
	c.Locals("email", email)
	return c.Next()
}
