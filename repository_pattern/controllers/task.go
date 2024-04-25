package controllers

import (
	"repo_pattern/models"
	"repo_pattern/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService}
}

// @Summary Create a New Task
// @Description Create a new task
// @Tags Task
// @Router /tasks [post]
func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := tc.taskService.CreateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

// @Summary Update an existing task
// @Description Update an existing task
// @Tags Task
// @Router /tasks/{id} [put]
func (tc *TaskController) UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	task.ID = uint(id)

	if err := tc.taskService.UpdateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

// @Summary Delete a task
// @Description Delete a task by ID
// @Tags Task
// @Router /tasks/{id} [delete]
func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	if err := tc.taskService.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// @Summary Get all tasks
// @Description Get all tasks
// @Tags Task
// @Router /tasks [get]
func (tc *TaskController) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

// @Summary Get task by ID
// @Description Get a task by ID
// @Tags Task
// @Router /tasks/{id} [get]
func (tc *TaskController) GetTaskById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task ID"})
	}

	task, err := tc.taskService.GetTaskById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}
