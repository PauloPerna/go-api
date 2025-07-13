package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/stat"
)

type SpearmanRequest struct {
	X []float64 `json:"x" binding:"required"`
	Y []float64 `json:"y" binding:"required"`
}

func SpearmanHandler(c *gin.Context) {
	var req SpearmanRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	//Compute Correlation
	corr := stat.Correlation(req.X, req.Y, nil)

	c.JSON(http.StatusOK, gin.H{"spearman": corr})
}
