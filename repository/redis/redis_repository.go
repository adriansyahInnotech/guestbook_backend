package redis

import (
	"context"
	"guestbook_backend/db"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	ValidateAccess(cardNumber, deviceID string) (bool, error)
	SetCardDevices(cardNumber string, deviceIDs []string) error
	GetCardDevices(cardNumber string) ([]string, error)
	RemoveCardDevices(cardNumber string) error
}

type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

// NewredisRepository creates a new Redis-based OTP repository
func NewRedisRepository() RedisRepository {
	return &redisRepository{
		client: db.ClientRedis,
		ctx:    context.Background(),
	}
}

func (s *redisRepository) ValidateAccess(cardNumber, deviceID string) (bool, error) {
	key := "card:" + cardNumber + ":devices"
	allowed, err := s.client.SIsMember(s.ctx, key, deviceID).Result()
	if err != nil {
		return false, err
	}
	return allowed, nil
}

func (s *redisRepository) SetCardDevices(cardNumber string, deviceIDs []string) error {
	// Hapus dulu
	s.client.Del(s.ctx, "card:"+cardNumber+":devices")
	if len(deviceIDs) > 0 {
		return s.client.SAdd(s.ctx, "card:"+cardNumber+":devices", deviceIDs).Err()
	}

	return nil
}

func (s *redisRepository) GetCardDevices(cardNumber string) ([]string, error) {
	return s.client.SMembers(s.ctx, "card:"+cardNumber+":devices").Result()
}

func (s *redisRepository) RemoveCardDevices(cardNumber string) error {
	return s.client.Del(s.ctx, "card:"+cardNumber+":devices").Err()
}
