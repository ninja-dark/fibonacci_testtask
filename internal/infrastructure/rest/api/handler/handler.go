package handler

import (

	"github.com/gin-gonic/gin"
	fibologic "github.com/ninja-dark/fibonacci_testtask/internal/fiboLogic"
)

type Handler struct {
	Services *fibologic.Fibo
}

func (h *Handler) InitRouters() *gin.Engine{
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/fibonacci", h.GetFibonacci)
	return router
}