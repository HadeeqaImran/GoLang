package main

import (
	"fmt"

	_ "repo_pattern/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @Summary Get a personalized greeting
// @Description Get a personalized greeting based on the provided name
// @ID get-greeting
// @Param name path string true "Name for the greeting"
// @Produce plain
// @Success 200 {string} string "Hello {name}"
// @Failure 400 {string} string "Bad Request"
// @Router /hello/{name} [get]
func hello(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.SendString(fmt.Sprintf("Hello, %s", name))
}

// @title Hello
func main() {
	// Creating a fiber app
	app := fiber.New()

	// Creating the Http Handler
	// :name here is a path parameter
	app.Get("hello/:name", hello)

	// Integrating Swagger for Using the API
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Listen for Requests
	app.Listen(":3000")
}
