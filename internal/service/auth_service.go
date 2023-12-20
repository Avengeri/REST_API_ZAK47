package service

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository"
	"fmt"
	"log"
)

type AuthService struct {
	repo repository.StorageUsers
}

func NewAuthService(repo repository.StorageUsers) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) Set(user *model.User) error {
	err := r.repo.Set(user)
	if err != nil {
		log.Printf("Ошибка при добавлении пользователя: %v", err)
	}
	return err
}
func (r *AuthService) Get(id int64) (*model.User, error) {
	user, err := r.repo.Get(id)
	if err != nil {
		log.Printf("Ошибка при получении пользователя: %v", err)
		return nil, err
	}
	return user, err
}
func (r *AuthService) Check(id int64) (bool, error) {
	exists, err := r.repo.Check(id)
	if err != nil {
		log.Printf("Ошибка при проверке пользователя: %v", err)
		return false, err
	}

	return exists, nil

}
func (r *AuthService) Delete(id int64) error {

	exists, err := r.repo.Check(id)
	if err != nil {
		return fmt.Errorf("ошибка при проверке пользователя: %v", err)
	}

	if !exists {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}

	if err = r.repo.Delete(id); err != nil {
		return fmt.Errorf("ошибка при удалении пользователя: %v", err)
	}
	return nil
}
func (r *AuthService) GetAllId() ([]int64, error) {
	ids, err := r.repo.GetAllId()
	if err != nil {
		// Обработка ошибки, например, логирование или возврат пустого списка и ошибки
		log.Printf("Ошибка при получении ID: %v", err)
		return nil, err
	}
	return ids, nil
}
