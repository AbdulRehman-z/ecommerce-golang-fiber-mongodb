package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"jwt-golang/services"
	"log"
)

const (
	port = ":3000"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	// routes

	// connect to mongodb
	services.ConnectWithMongodb()
	// listen to the port + start the server
	if err := app.Listen(port); err != nil {
		log.Fatal("Error starting the server: ", err.Error())
	}
}
