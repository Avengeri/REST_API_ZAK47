package service

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository"
	"fmt"
	"log"
)

type TodoService struct {
	repo repository.StorageUsers
}

func NewTodoService(repo repository.StorageUsers) *TodoService {
	return &TodoService{repo: repo}
}

func (r *TodoService) Set(user *model.User) error {
	err := r.repo.Set(user)
	if err != nil {
		log.Printf("error when adding a user: %v", err)
	}
	return err
}
func (r *TodoService) GetById(id int64) (*model.User, error) {
	user, err := r.repo.GetById(id)
	if err != nil {
		log.Printf("error receiving the user: %v", err)
		return nil, err
	}
	return user, err
}
func (r *TodoService) CheckById(id int64) (bool, error) {
	exists, err := r.repo.CheckById(id)
	if err != nil {
		log.Printf("error checking the user: %v", err)
		return false, err
	}

	return exists, nil

}

func (r *TodoService) Delete(id int64) error {
	exists, err := r.CheckById(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the user has been not found")
	}
	if err = r.repo.Delete(id); err != nil {
		return fmt.Errorf("error when deleting a user: %v", err)
	}
	return nil
}
func (r *TodoService) GetAllId() ([]int64, error) {
	ids, err := r.repo.GetAllId()
	if err != nil {
		log.Printf("error receiving the ID: %v", err)
		return nil, err
	}
	return ids, nil
}
