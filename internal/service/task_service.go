package service

import (
	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

// Create a new task
func (s *TaskService) Create(task *model.Task) error {
	// Additional validation or business logic could go here
	return s.repo.Create(task)
}

// GetAll returns all tasks (optionally filtered by userID)
func (s *TaskService) GetAll(userID *uint) ([]model.Task, error) {
	return s.repo.GetAll(userID)
}

// Update an existing task
func (s *TaskService) Update(task *model.Task) error {
	return s.repo.Update(task)
}

// Delete a task (optionally enforcing user ownership)
func (s *TaskService) Delete(id uint, userID *uint) error {
	return
