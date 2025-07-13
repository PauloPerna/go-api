package main

import (
	"net/http"

	"gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/spearman", handlers.SpearmanHandler)

	router.POST("/linreg", handlers.LinregHandler)

	router.Run() // listen and serve on 0.0.0.0:8080
}
