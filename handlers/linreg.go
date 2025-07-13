package handlers

import (
	"gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/mat"
)

func LinregHandler(c *gin.Context) {
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
	var Xt mat.Dense
	var XtX mat.Dense
	var inv mat.Dense
	var invXt mat.Dense
	var beta mat.Dense

	X, _ := utils.ToDense(raw, true) // My function to parse raw into *mat.Dense
	Xt.CloneFrom(X.T())
	y_vec := utils.ToVecDense(y) // My function to parse raw into *mat.VecDense

	XtX.Mul(&Xt, X)
	inv.Inverse(&XtX)
	invXt.Mul(&inv, &Xt)
	beta.Mul(&invXt, y_vec)

	// Extract the coefficients as a slice
	rows, cols := beta.Dims()
	coeffs := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		coeffs[i] = beta.At(i, 0)
	}

	// Return Model Parameters
	c.JSON(http.StatusOK, gin.H{
		"intercept":    coeffs[0],
		"coefficients": coeffs[1:],
	})
}
