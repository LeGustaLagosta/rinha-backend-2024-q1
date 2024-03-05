package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"rinha/internal/model"
	"rinha/internal/repository"
)

func getCliente(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "id inválido"})
		return
	}

	var transacao model.Transacao
	if err := c.BindJSON(&transacao); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "transação inválida"})
        return
    }

	transacao.ID_cliente = id
	
	err = repository.InserirTransacao(&transacao)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"mensagem": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"mensagem": "transação registrada"})
}
