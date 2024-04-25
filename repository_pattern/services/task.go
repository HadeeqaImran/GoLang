package services

import (
	"repo_pattern/models"
	"repo_pattern/repositories"
)

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepo}
}

func (ts *TaskService) CreateTask(task *models.Task) error {
	// Perform any business logic validations here
	return ts.taskRepository.Create(task)
}

func (ts *TaskService) UpdateTask(task *models.Task) error {
	// Perform any business logic validations here
	return ts.taskRepository.Update(task)
}

func (ts *TaskService) DeleteTask(id uint) error {
	// Perform any business logic validations here
	return ts.taskRepository.Delete(id)
}

func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	return ts.taskRepository.GetAll()
}

func (ts *TaskService) GetTaskById(id uint) (*models.Task, error) {
	return ts.taskRepository.GetById(id)
}
