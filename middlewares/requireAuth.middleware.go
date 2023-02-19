package middlewares

import (
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
	id, email, userType, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token",
		})
	}

	// set user id and email to context
	c.Locals("id", id)
	c.Locals("email", email)
	c.Locals("userType", userType)
	return c.Next()
}
