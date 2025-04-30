package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	"auth-system/app/domain"
	"gorm.io/gorm/schema"
)

func InitDB(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
	})
    if err != nil {
        log.Fatal("Falha na conexão ao banco de dados:", err)
    }
	db.AutoMigrate(
		&domain.User{},
		&domain.Consumidor{},
		&domain.Empresa{},
		&domain.Divida{},
		&domain.NaturezaCobranca{},
		&domain.DividaConsumidorEmpresa{},
	)
    return db
}
