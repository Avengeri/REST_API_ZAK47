package repository

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository/postgres"
	"Interface_droch_3/internal/repository/redis_storage"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type StorageUsers interface {
	Set(user *model.User) error
	GetById(id int64) (*model.User, error)
	CheckById(id int64) (bool, error)
	Delete(id int64) error
	GetAllId() ([]int64, error)
}

type Repository struct {
	StorageUsers
	Authorization
}

func NewStorageUsersPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		StorageUsers:  postgres.NewTodoPostgres(db),
		Authorization: postgres.NewAuthPostgres(db),
	}
}

func NewStorageUsersRedis(rdb *redis.Client) *Repository {
	return &Repository{
		StorageUsers: redis_storage.NewAuthRedis(rdb)}
}
