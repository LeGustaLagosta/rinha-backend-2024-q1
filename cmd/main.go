package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"database/sql"
	_ "github.com/lib/pq"

	"rinha/internal/api"
	"rinha/internal/repository"
)

func main() {
	db_hostname := os.Getenv("DB_HOSTNAME")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_hostname, db_user, db_password, db_database, db_port)
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("falha ao conectar com o banco de dados")
	}
	repository.InitDB(DB)
	defer repository.CloseDB()

	http.HandleFunc("/clientes/", api.RouteClientes)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
