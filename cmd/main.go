package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"rinha/internal/api"
	"rinha/internal/repository"
)

func main() {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := "host=localhost user=admin password=123 dbname=rinha port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar com o banco de dados")
	}
	repository.InitDB(DB)

	router := api.GetRouter()
	router.Run(":5001")
}
