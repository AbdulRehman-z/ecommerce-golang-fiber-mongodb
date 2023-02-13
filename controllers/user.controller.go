package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
)

var collection *mongo.Collection = database.OpenCollection(database.Client, "users")

func Signup(c *fiber.Ctx) error {

}

func Signin(c *fiber.Ctx) error {

}

func Logout(c *fiber.Ctx) error {

}

func CurrentUser(c *fiber.Ctx) error {

}
