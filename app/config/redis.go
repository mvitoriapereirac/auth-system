package config

import (
    "context"
    "log"
    "github.com/go-redis/redis/v8"
)

var (
    Ctx = context.Background()
    RDB *redis.Client
)

func InitRedis() *redis.Client {
    RDB = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })

    if _, err := RDB.Ping(Ctx).Result(); err != nil {
        log.Fatal("Falha na conexão com Redis:", err)
    }

	return RDB;
}
