package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/stat"
)

type SpearmanRequest struct {
	X []float64 `json:"x" binding:"required"`
	Y []float64 `json:"y" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/spearman", func(c *gin.Context) {
		var req SpearmanRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		//Compute Correlation
		corr := stat.Correlation(req.X, req.Y, nil)

		c.JSON(http.StatusOK, gin.H{"spearman": corr})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
