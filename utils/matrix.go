package utils

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func ToDense(X map[string][]float64, addIntercept bool) (*mat.Dense, error) {
	// Enforce deterministic column order
	keys := make([]string, 0, len(X))
	for key := range X {
		keys = append(keys, key)
	}

	// Check length
	var rows int
	for i, key := range keys {
		if i == 0 {
			rows = len(X[key])
		} else {
			if len(X[key]) != rows {
				return nil, fmt.Errorf("column %s has inconsistent length", key)
			}
		}
	}

	cols := len(keys)
	if addIntercept {
		cols++
	}

	flatData := make([]float64, 0, rows*cols)

	for i := 0; i < rows; i++ {
		if addIntercept {
			flatData = append(flatData, 1.0)
		}
		for _, key := range keys {
			flatData = append(flatData, X[key][i])
		}
	}

	return mat.NewDense(rows, cols, flatData), nil
}

func ToVecDense(y []float64) *mat.VecDense {
	return mat.NewVecDense(len(y), y)
}
