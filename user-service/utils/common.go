package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Println(message)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}

func ResponseWithData(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
