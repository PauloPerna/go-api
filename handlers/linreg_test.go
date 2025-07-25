package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLinregHandler(t *testing.T) {
	// Define the request body (we expect intercept = 1, coef = [1])
	requestBody := map[string][]float64{
		"y": {2, 3, 4},
		"x": {1, 2, 3},
	}
	body, _ := json.Marshal(requestBody)

	// Create Request
	req, err := http.NewRequest(http.MethodPost, "/linreg", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-type", "application/json")

	// Create a Response Recorder
	w := httptest.NewRecorder()

	// Create a Gin context from the request and recorder
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the Handler
	LinregHandler(c)

	// Check Status
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", w.Code)
	}

	// Define a struct to unmarshal the JSON response
	var resp struct {
		Intercept    float64   `json:"intercept"`
		Coefficients []float64 `json:"coefficients"`
	}

	// Parse response body

	// Assert intercept and coefficients

}
