package main

import (
    "fmt"

	"rinha/internal/api"
	"rinha/internal/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	clientes := carregarCadastroInicial("cadastro-inicial.json")
    fmt.Println("cadastro-inicial:", clientes)
	DB, err := gorm.Open(sqlite.Open("rinha.db"), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar com o banco de dados")
	}
	repository.InitDB(DB)
	repository.PersistirCadastroInicial(clientes)

	router := api.GetRouter()
	router.Run(":5001")
}
