package routes

import (
	"github.com/gofiber/fiber/v2"
	"jwt-golang/controllers"
	"jwt-golang/middlewares"
)

func Router(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", controllers.Welcome)

	//auth routes
	userApi := api.Group("/users/auth")
	userApi.Post("/signup", middlewares.ValidateCredentialsMiddleware, controllers.Signup)
	userApi.Post("/signin", middlewares.ValidateCredentialsMiddleware, controllers.Signin)
	userApi.Post("/signout", middlewares.RequireAuthMiddleware, controllers.Signout)
	userApi.Get("/profile", middlewares.RequireAuthMiddleware, controllers.Profile)

	// products routes
	productsApi := api.Group("/products")
	productsApi.Get("/", controllers.GetAllProducts)
	productsApi.Get("/:id", controllers.GetProduct)
	productsApi.Post("/create", middlewares.RequireAuthMiddleware, controllers.CreateProduct)
	productsApi.Put("/:id", middlewares.RequireAuthMiddleware, controllers.UpdateProduct)
	productsApi.Delete("/:id", middlewares.RequireAuthMiddleware, controllers.DeleteProduct)

	// cart routes
	cartApi := api.Group("/cart", middlewares.RequireAuthMiddleware)
	cartApi.Post("/remove/:id", controllers.RemoveProductFromCart)
	cartApi.Post("/add/:id", controllers.AddProductToCart)

	// address routes
	addressApi := api.Group("/address", middlewares.RequireAuthMiddleware)
	addressApi.Put("/update", controllers.UpdateAddress)

	// order routes
	orderApi := api.Group("/order", middlewares.RequireAuthMiddleware)
	orderApi.Post("/", controllers.OrderAll)
	orderApi.Post("/:id", controllers.OrderOne)

	// admin routes
	adminApi := api.Group("/admin", middlewares.RequireAuthMiddleware)
	adminApi.Get("/getUser/:id", controllers.GetUser)
	adminApi.Get("/getUsers", controllers.GetUsers)
	adminApi.Delete("/deleteUser/:id", controllers.DeleteUser)
	adminApi.Delete("/deleteUsers", controllers.DeleteAllUsers)

}
