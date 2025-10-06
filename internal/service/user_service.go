package service

import (
	"errors"

	"github.com/g-laliotis/task-tracker/internal/model"
	"github.com/g-laliotis/task-tracker/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(user *model.User) error {
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Create(user)
}

func (s *UserService) Authenticate(email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
