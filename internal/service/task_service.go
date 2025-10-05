package service

import (
	"errors"

	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) CreateTask(title string) (*model.Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	task := &model.Task{Title: title}
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(id uint, completed bool) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	task.Completed = completed
	return s.repo.Update(&task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
