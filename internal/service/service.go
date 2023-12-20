package service

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository"
)

type ServiceUsers interface {
	Set(user *model.User) error
	Get(id int64) (*model.User, error)
	Check(id int64) (bool, error)
	Delete(id int64) error
	GetAllId() ([]int64, error)
}

type Service struct {
	ServiceUsers
}

func NewServiceUsers(repos *repository.Repository) *Service {
	return &Service{ServiceUsers: repos.StorageUsers}
}
