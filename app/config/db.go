package config

import (
	"auth-system/app/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
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

	insertDefaultNaturezaCobranca(db) //Idealmente, aqui haveria um arquivo de migrations
	return db
}

func insertDefaultNaturezaCobranca(db *gorm.DB) {
	var count int64
	db.Model(&domain.NaturezaCobranca{}).Count(&count)

	if count == 0 {
		defaultData := []domain.NaturezaCobranca{
			{RazaoCobranca: "Cobrança Simples"},
			{RazaoCobranca: "Cobrança Judicial"},
			{RazaoCobranca: "Negociação Amigável"},
		}

		result := db.Create(&defaultData)
		if result.Error != nil {
			log.Println("Erro ao inserir dados padrão:", result.Error)
		} else {
			log.Println("Dados padrão inseridos")
		}
	}
}
