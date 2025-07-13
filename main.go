package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/stat"
	//	"gonum.org/v1/gonum/mat" < will be used to evaluate OLS
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

	router.POST("/linreg", func(c *gin.Context) {
		var raw map[string][]float64

		// Parse requested body
		if err := c.BindJSON(&raw); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON Body"})
			return
		}

		// Validate presence of y
		y, ok := raw["y"]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You must pass a variable named \"y\""})
			return
		}
		delete(raw, "y") // Remove y from independent variables

		// Validate presence of independent variables
		if len(raw) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You must pass independent variables"})
		}

		// Validate all arrays has the same length
		n := len(y)
		for name, col := range raw {
			if len(col) != n {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "all arrays must have the same length",
					"field": name,
				})
				return
			}
		}

		// Evaluate OLS

		// Return Model Parameters
		c.JSON(http.StatusOK, gin.H{
			"intercept":    nil,
			"coefficients": nil,
		})

	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
