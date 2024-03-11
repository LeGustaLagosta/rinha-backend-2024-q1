package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"rinha/internal/model"
	"rinha/internal/repository"
)

type Mensagem_Erro struct {
	Mensagem string `json:"mensagem"`
}

func RouteClientes(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	part := pathParts[3]

	switch part {
	case "extrato":
		GetExtrato(w, r)
	case "transacoes":
		PostTransacao(w, r)
	default:
		http.Error(w, "Invalid endpoint", http.StatusNotFound)
	}
}

func GetExtrato(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(extractID(r.URL.Path), 10, 64)
	if err != nil {
		sendJSONResponse(w, 404, Mensagem_Erro{Mensagem: "cliente não encontrado: id inválido"})
		return
	}

	cliente, err := repository.ObterCliente(id)
	if err != nil {
		sendJSONResponse(w, 404, Mensagem_Erro{Mensagem: "cliente não encontrado: " + err.Error()})
		return
	}

	transacoes, err := repository.ObterTransacoes(id)
	if err != nil {
		sendJSONResponse(w, 404, Mensagem_Erro{Mensagem: "transações não encontradas: " + err.Error()})
		return
	}

	saldo := &model.Saldo{
		Total: cliente.Saldo,
		Data_extrato: time.Now().UTC(),
		Limite: cliente.Limite,
	}

	extrato := &model.Extrato{
		Saldo_Cliente: saldo,
		Ultimas_Transacoes: transacoes,
	}
	sendJSONResponse(w, 200, extrato)
}

func PostTransacao(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(extractID(r.URL.Path), 10, 64)
	if err != nil {
		sendJSONResponse(w, 404, Mensagem_Erro{Mensagem: "cliente não encontrado: id inválido"})
		return
	}

	var transacao model.Transacao
	if err := json.NewDecoder(r.Body).Decode(&transacao); err != nil {
		sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "transação inválida: " + err.Error()})
		return
	}

	// validar regras
	if transacao.Valor < 0 {
		sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "transação inválida: valor não pode ser negativo"})
		return
	}

	if transacao.Tipo != "c" && transacao.Tipo != "d" {
		sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "transação inválida: tipo diferente de c ou d"})
		return
	}

	if utf8.RuneCountInString(transacao.Descricao) > 10 {
		sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "transação inválida: descrição maior do que 10 caracteres"})
		return
	}

	transacao.ID_cliente = id
	cliente, err := repository.ObterCliente(id)
	if err != nil {
		sendJSONResponse(w, 404, Mensagem_Erro{Mensagem: "cliente não encontrado: " + err.Error()})
		return
	}

	var novoSaldo int64
	if transacao.Tipo == "d" {
		novoSaldo = cliente.Saldo - transacao.Valor
		if novoSaldo < (cliente.Limite * -1) {
			sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "saldo não pode superar o limite"})
			return
		}
	} else if transacao.Tipo == "c" {
		novoSaldo = cliente.Saldo + transacao.Valor
	}

	cliente.Saldo = novoSaldo
	transacao.Data = time.Now()
	
	err = repository.InserirTransacao(&transacao, cliente)
	if err != nil {
		sendJSONResponse(w, 422, Mensagem_Erro{Mensagem: "transação não registrada: " + err.Error()})
		return
	}

	sendJSONResponse(w, 200, cliente)
}

func extractID(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}