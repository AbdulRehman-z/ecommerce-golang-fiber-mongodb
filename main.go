package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	services "jwt-golang/services"
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

	// connect to mongodb
	services.ConnectWithMongodb()
	// listen to the port + start the server
	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}
