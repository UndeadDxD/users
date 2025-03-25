package service

import (
	"users/internal/models"
	"users/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}
type userService struct {
	repo repository.UserRepository
}

func (s userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s userService) GetById(id int) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
