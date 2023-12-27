package redis_storage

import (
	"Interface_droch_3/internal/model"
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

type AuthRedis struct {
	rdb *redis.Client
}

func NewAuthRedis(rdb *redis.Client) *AuthRedis {
	return &AuthRedis{rdb: rdb}
}

func (r *AuthRedis) Set(user *model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := "user:" + strconv.Itoa(int(user.Id))
	return r.rdb.Set(key, userJSON, 5*time.Minute).Err()
}

func (r *AuthRedis) GetById(id int64) (*model.User, error) {
	key := "user:" + strconv.Itoa(int(id))
	userJSON, err := r.rdb.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRedis) CheckById(id int64) (bool, error) {
	key := "user:" + strconv.Itoa(int(id))

	exists, err := r.rdb.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r *AuthRedis) Delete(id int64) error {
	key := "user:" + strconv.Itoa(int(id))
	_, err := r.rdb.Del(key).Result()
	return err
}

func (r *AuthRedis) GetAllId() ([]int64, error) {
	keys, err := r.rdb.Keys("user:*").Result()
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) >= 2 {
			idStr := parts[1]
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err == nil {
				ids = append(ids, id)
			}
		}
	}
	return ids, nil
}
