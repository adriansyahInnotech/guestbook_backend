package db

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ClientRedis *redis.Client
var Ctx context.Context = context.Background()

// InitRedis initializes Redis client
func InitRedis() {
	ClientRedis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // jika pakai password isi di sini
		DB:       0,
	})

	// Ping untuk test koneksi
	_, err := ClientRedis.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("❌ Redis connection error: %v", err))
	}
	fmt.Println("✅ Redis connected successfully")
}
