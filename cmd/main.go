package main

import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"rinha/internal/api"
	"rinha/internal/repository"
)

func main() {
	db_hostname := os.Getenv("DB_HOSTNAME")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_hostname, db_user, db_password, db_database, db_port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar com o banco de dados")
	}
	repository.InitDB(DB)

	router := api.GetRouter()
	router.Run(":8080")
}
