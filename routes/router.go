package routes

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", controller.Welcome)

	// auth routes
	authApi := api.Group("/auth")
	authApi.Post("/signup", controller.Register)
	authApi.Post("/login", controller.Login)
	authApi.Post("/logout", controller.Logout)

	// user routes
	userApi := api.Group("/user")
	userApi.Get("/currentUser", controller.CurrentUser)

	// products routes
	productsApi := api.Group("/products")
	productsApi.Get("/", controller.GetAllProducts)
	productsApi.Get("/:id", controller.GetProduct)
	productsApi.Post("/", controller.CreateProduct)
	productsApi.Put("/:id", controller.UpdateProduct)
	productsApi.Delete("/:id", controller.DeleteProduct)

	// admin routes
	adminApi := api.Group("/admin")
	adminApi.Get("/allUsers", controller.GetAllUsers)
	adminApi.Delete("/deleteUser/:id", controller.DeleteUser)
	adminApi.Delete("/deleteAllUsers", controller.DeleteAllUsers)
}
