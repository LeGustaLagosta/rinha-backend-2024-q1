package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"rinha/internal/model"
	"rinha/internal/repository"
)

func getExtrato(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado"})
		return
	}

	cliente, err := repository.ObterCliente(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado"})
		return
	}

	transacoes, err := repository.ObterTransacoes(id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"saldo": gin.H{
			"total": cliente.Saldo,
			"data_extrato": time.Now().UTC(),
			"limite": cliente.Limite,
		},
		"ultimas_transacoes": transacoes,
	})
}

func postTransacao(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado"})
		return
	}

	var transacao model.Transacao
	if err := c.ShouldBindJSON(&transacao); err != nil {
		c.IndentedJSON(422, gin.H{"mensagem": "transação inválida: " + err.Error()})
        return
    }

	transacao.ID_cliente = id
	cliente, err := repository.ObterCliente(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado"})
		return
	}

	var novoSaldo int64
	if transacao.Tipo == "d" {
		novoSaldo = cliente.Saldo - transacao.Valor
		if novoSaldo < (cliente.Limite * -1) {
			c.IndentedJSON(422, gin.H{"mensagem": "saldo não pode superar o limite"})
			return
		}
	} else if transacao.Tipo == "c" {
		novoSaldo = cliente.Saldo + transacao.Valor
	}

	cliente.Saldo = novoSaldo
	transacao.Data = time.Now()
	
	err = repository.InserirTransacao(&transacao, cliente)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"mensagem": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensagem": "transação registrada"})
}
