package repository

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"guestbook_backend/db"
// 	"guestbook_backend/models"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// )

// type RedisRepository interface {
// 	SetPendingUser(email string, pendingUser *models.RedisPendingUserModel, expiration time.Duration) error
// 	GetPendingUser(email string, otp string) (*models.RedisPendingUserModel, error)
// 	DeletePendingUser(email string) error
// 	SetForgotUser(email string, forgotUser *models.RedisForgotUserModel, expiration time.Duration) error
// 	GetForgotUser(id string) error
// 	DeleteForgotUser(id string) error
// }

// type redisRepository struct {
// 	client *redis.Client
// 	ctx    context.Context
// }

// // NewredisRepository creates a new Redis-based OTP repository
// func NewRedisRepository() RedisRepository {
// 	return &redisRepository{
// 		client: db.ClientRedis,
// 		ctx:    context.Background(),
// 	}
// }

// func (r *redisRepository) SetPendingUser(email string, pendingUser *models.RedisPendingUserModel, expiration time.Duration) error {
// 	key := fmt.Sprintf("pending_user:%s-%s", email, pendingUser.Otp)

// 	data, err := json.Marshal(pendingUser)
// 	if err != nil {
// 		return err
// 	}

// 	return r.client.Set(r.ctx, key, data, expiration).Err()
// }

// func (r *redisRepository) GetPendingUser(email string, otp string) (*models.RedisPendingUserModel, error) {
// 	pendingUser := new(models.RedisPendingUserModel)

// 	key := fmt.Sprintf("pending_user:%s-%s", email, otp)

// 	result, err := r.client.Get(r.ctx, key).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := json.Unmarshal([]byte(result), pendingUser); err != nil {
// 		return nil, err
// 	}

// 	return pendingUser, nil
// }

// func (r *redisRepository) DeletePendingUser(email string) error {
// 	key := fmt.Sprintf("pending_user:%s", email)
// 	return r.client.Del(r.ctx, key).Err()
// }

// func (r *redisRepository) SetForgotUser(id string, forgotUser *models.RedisForgotUserModel, expiration time.Duration) error {
// 	key := fmt.Sprintf("forgot_user:%s", id)

// 	data, err := json.Marshal(forgotUser)
// 	if err != nil {
// 		return err
// 	}

// 	return r.client.Set(r.ctx, key, data, expiration).Err()
// }

// func (r *redisRepository) GetForgotUser(id string) error {

// 	key := fmt.Sprintf("forgot_user:%s", id)

// 	_, err := r.client.Get(r.ctx, key).Result()
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func (r *redisRepository) DeleteForgotUser(id string) error {
// 	key := fmt.Sprintf("forgot_user:%s", id)
// 	return r.client.Del(r.ctx, key).Err()
// }
