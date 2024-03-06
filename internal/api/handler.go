package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"rinha/internal/model"
	"rinha/internal/repository"
)

func getCliente(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "id inválido"})
		return
	}

	cliente, err := repository.ObterCliente(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"mensagem": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, cliente)
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
	
	err = repository.InserirTransacao(&transacao, cliente)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"mensagem": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensagem": "transação registrada"})
}
