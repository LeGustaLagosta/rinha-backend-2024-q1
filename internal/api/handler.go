package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

}
