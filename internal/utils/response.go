package utils

import (
	"be/internal/models"
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data interface{}, status string) {
	response := models.CustomResponse{
		Data:   data,
		Status: status,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
