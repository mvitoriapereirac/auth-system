package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func InitDB(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha na conexão ao banco de dados:", err)
    }
    return db
}
