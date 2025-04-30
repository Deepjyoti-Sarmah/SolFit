package handlers

import (
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ok",
		Message:   "Server is running",
		Version:   "0.1.0",
		Timestamp: time.Now(),
	}

	RespondWithSuccess(w, http.StatusOK, "Health check successful", response)
}
