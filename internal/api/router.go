package api

import (
	"github.com/gin-gonic/gin"
)

func newCustomGinEngine() *gin.Engine {
	// Create a new Gin engine with a custom configuration
	r := gin.New()

	// Disable console color
	gin.DisableConsoleColor()

	// Set release mode for better performance
	gin.SetMode(gin.ReleaseMode)

	// // Use Gzip middleware for response compression
	// r.Use(gin.Recovery())
	// r.Use(gin.Logger())
	// r.Use(gin.ErrorLogger())

	return r
}

func GetRouter() *gin.Engine {
	// router := gin.Default()
	router := newCustomGinEngine()
	router.GET("/clientes/:id/extrato", getExtrato)
	router.POST("/clientes/:id/transacoes", postTransacao)

	return router
}
