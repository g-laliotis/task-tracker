package repository

import (
	"github.com/g-laliotis/task-tracker/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetAll(userID uint) ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id uint, userID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAll(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Task{}).Error
}
