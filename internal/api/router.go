package api

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/clientes/:id/extrato", getCliente)
	router.POST("/clientes/:id/transacoes", postTransacao)

	return router
}
