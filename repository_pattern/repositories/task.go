package repositories

import (
	"errors"
	"repo_pattern/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) Create(task *models.Task) error {
	result := tr.db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) Update(task *models.Task) error {
	result := tr.db.Model(&models.Task{}).Where("id = ?", task.ID).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no task found with the given ID")
	}
	return nil
}

func (tr *TaskRepository) Delete(id uint) error {
	result := tr.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no task found with the given ID")
	}
	return nil
}

func (tr *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := tr.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tr *TaskRepository) GetById(id uint) (*models.Task, error) {
	var task models.Task
	result := tr.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}
