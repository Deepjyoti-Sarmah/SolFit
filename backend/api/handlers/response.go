package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ReponseWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, Response{
		Success: false,
		Error:   message,
	})
}

func RespondWithSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	RespondWithJSON(w, code, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
