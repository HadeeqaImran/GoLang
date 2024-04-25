package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

// @Summary Get a personalized greeting
// @Description Get a personalized greeting based on the provided name
// @ID get-greeting
// @Param name path string true "Name for the greeting"
// @Produce plain
// @Success 200 {string} string "Hello {name}"
// @Failure 400 {string} string "Bad Request"
// @Router /hello/{name} [get]
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcom to my app")
}

// Setup Endpoints
func setupRoutes(app *fiber.App) {
	// welcome page
	app.Get("/api", welcome)
	// User Endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Product Endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("api/products/:id", routes.DeleteProduct)
	// Order Endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetAllOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Delete("/api/orders/:id", routes.DeleteOrder)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	// Integrating Swagger for Using the API
	app.Get("/swagger/*", swagger.HandlerDefault)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
