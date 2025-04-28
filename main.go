package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    DB    *gorm.DB
    RDB   *redis.Client
    ctx   = context.Background()
)

func initPostgres() {
    dsn := "host=db user=postgres password=postgres dbname=auth_system port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to Postgres:", err)
    }
    fmt.Println("Connected to Postgres")
}

func initRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })
    if _, err := RDB.Ping(ctx).Result(); err != nil {
        log.Fatal("Failed to connect to Redis:", err)
    }
    fmt.Println("Connected to Redis")
}

func main() {
    initPostgres()
    initRedis()

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    r.Run(":8080")
}