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

func (s *TaskService) Create(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) Update(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id uint) error {
	return s.repo.Delete(id)
}
