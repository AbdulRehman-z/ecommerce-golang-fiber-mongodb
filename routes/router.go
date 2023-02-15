package routes

import (
	"github.com/gofiber/fiber/v2"
	"jwt-golang/controllers"
)

func Router(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", controllers.Welcome)

	//auth routes
	userApi := api.Group("/users/auth")
	userApi.Post("/signup", controllers.Signup)
	userApi.Post("/login", controllers.Signin)
	userApi.Post("/logout", controllers.Logout)
	userApi.Get("/profile", controllers.Profile)

	// user routes

	// products routes
	productsApi := api.Group("/products")
	productsApi.Get("/", controllers.GetAllProducts)
	productsApi.Get("/:id", controllers.GetProduct)
	productsApi.Post("/", controllers.CreateProduct)
	productsApi.Put("/:id", controllers.UpdateProduct)
	productsApi.Delete("/:id", controllers.DeleteProduct)

	// admin routes
	adminApi := api.Group("/admin")
	adminApi.Get("/getUser/:id", controllers.GetUser)
	adminApi.Get("/allUsers", controllers.GetAllUsers)
	adminApi.Delete("/deleteUser/:id", controllers.DeleteUser)
	adminApi.Delete("/deleteAllUsers", controllers.DeleteAllUsers)
}
