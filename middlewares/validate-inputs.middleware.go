package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateCredentialsMiddleware(c *fiber.Ctx) error {
	validate := validator.New()

	type UserInputCred struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// get the user inputs
	var user UserInputCred
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "missing field",
		})
	}

	// validate the user inputs
	if err := validate.Struct(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"err":     err.Error(),
		})
	}

	return c.Next()
}
