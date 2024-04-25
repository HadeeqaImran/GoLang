package main

import (
	"log"

	"repo_pattern/controllers"
	"repo_pattern/database"
	_ "repo_pattern/docs"
	"repo_pattern/repositories"
	"repo_pattern/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func setupRoutes(app *fiber.App, taskController *controllers.TaskController) {
	// Routes for tasks
	app.Post("/tasks", taskController.CreateTask)
	app.Put("/tasks/:id", taskController.UpdateTask)
	app.Delete("/tasks/:id", taskController.DeleteTask)
	app.Get("/tasks", taskController.GetAllTasks)
	app.Get("/tasks/:id", taskController.GetTaskById)
}

func main() {
	database.ConnectDb()

	// Initialize repository
	taskRepo := repositories.NewTaskRepository(database.Database.Db)

	// Initialize services
	taskService := services.NewTaskService(taskRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskService)

	app := fiber.New()

	// Integrating Swagger for Using the API
	app.Get("/swagger/*", swagger.HandlerDefault)
	setupRoutes(app, taskController)

	log.Fatal(app.Listen(":3000"))
}
