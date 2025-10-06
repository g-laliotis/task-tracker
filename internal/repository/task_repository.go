package repository

import (
	"github.com/g-laliotis/task-tracker/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetAll(userID *uint) ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id uint, userID *uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Create a new task
func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

// GetAll returns all tasks (optionally filtered by userID)
func (r *taskRepository) GetAll(userID *uint) ([]model.Task, error) {
	var tasks []model.Task
	query := r.db

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.Find(&tasks).Error
	return tasks, err
}

// Update an existing task
func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

// Delete a task by ID (and userID if provided)
func (r *taskRepository) Delete(id uint, userID *uint) error {
	query := r.db.Model(&model.Task{}).Where("id = ?", id)
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	return query.Delete(&model.Task{}).Error
}
