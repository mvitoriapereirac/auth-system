package main

import (
	"auth-system/app/config"
	"auth-system/app/interfaces"
)

func main() {
	dsn := "host=db user=postgres password=postgres dbname=auth_system port=5432 sslmode=disable"

	db := config.InitDB(dsn)
	redisClient := config.InitRedis()

	server := interfaces.NewServer(db, redisClient)
	server.Run(":8080")

}
