package repository

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository/postgres"
	"Interface_droch_3/internal/repository/redis_storage"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type StorageUsers interface {
	Set(user *model.User) error
	Get(id int64) (*model.User, error)
	Check(id int64) (bool, error)
	Delete(id int64) error
	GetAllId() ([]int64, error)
}

type Repository struct {
	StorageUsers
}

func NewStorageUsersPostgres(db *sqlx.DB) *Repository {
	return &Repository{StorageUsers: postgres.NewAuthPostgres(db)}
}

func NewStorageUsersRedis(rdb *redis.Client) *Repository {
	return &Repository{StorageUsers: redis_storage.NewAuthRedis(rdb)}
}
