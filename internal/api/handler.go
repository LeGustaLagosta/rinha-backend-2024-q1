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
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado: id inválido"})
		return
	}

	cliente, err := repository.ObterCliente(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado: " + err.Error()})
		return
	}

	transacoes, err := repository.ObterTransacoes(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "transações não encontradas: " + err.Error()})
		return
	}

	saldo := &model.Saldo{
		Total: cliente.Saldo,
		Data_extrato: time.Now().UTC(),
		Limite: cliente.Limite,
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"saldo": saldo,
		"ultimas_transacoes": transacoes,
	})
}

func postTransacao(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado: id inválido"})
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
		c.IndentedJSON(404, gin.H{"mensagem": "cliente não encontrado: " + err.Error()})
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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"mensagem": "transação não registrada: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensagem": "transação registrada"})
}
