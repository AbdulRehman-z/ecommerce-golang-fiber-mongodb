package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"jwt-golang/database"
	"jwt-golang/routes"
	"log"
)

const (
	port = ":3000"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	//store := session.New()

	// register session store middleware
	//app.Use(store)
	// register routes
	routes.Router(app)

	//routesStack := app.Stack()

	//for _, route := range routesStack {
	//	//log.Println(route.P)
	//}

	// connect to mongodb
	client := database.ConnectWithMongodb()

	// close the connection
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.Background())

	// listen to the port + start the server
	if err := app.Listen(port); err != nil {
		log.Fatal("Error starting the server: ", err.Error())
	}
}
