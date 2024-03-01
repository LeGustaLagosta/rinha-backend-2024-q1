package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rinha/internal/model"
)

func carregarCadastroInicial(json_filename string) []model.Cliente {
	// carga do cadastro inicial
	file, err := os.Open(json_filename)
	if err != nil {
		panic("falha ao abrir arquivo de cadastro inicial")
	}
	defer file.Close()

	var clientes []model.Cliente
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&clientes); err != nil {
		panic("falha ao decodificar arquivo de cadastro inicial")
	}
	fmt.Println("carga inicial: ", clientes)

	return clientes
}
