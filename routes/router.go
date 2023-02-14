package routes

import (
	"github.com/gofiber/fiber/v2"
	"jwt-golang/controllers"
)

func Router(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", controllers.Welcome)

	//auth routes
	authApi := api.Group("/auth")
	authApi.Post("/signup", controllers.Signup)
	authApi.Post("/login", controllers.Signin)
	authApi.Post("/logout", controllers.Logout)

	// user routes
	userApi := api.Group("/user")
	userApi.Get("/currentUser", controllers.CurrentUser)

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
