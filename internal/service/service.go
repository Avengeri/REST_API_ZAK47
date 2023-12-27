package service

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type ServiceUsers interface {
	Set(user *model.User) error
	GetById(id int64) (*model.User, error)
	CheckById(id int64) (bool, error)
	Delete(id int64) error
	GetAllId() ([]int64, error)
}

type Service struct {
	User     ServiceUsers
	AuthUser Authorization
}

func NewServiceUsers(repos *repository.Repository) *Service {
	return &Service{
		User:     NewTodoService(repos.StorageUsers),
		AuthUser: NewAuthService(repos.Authorization),
	}
}
