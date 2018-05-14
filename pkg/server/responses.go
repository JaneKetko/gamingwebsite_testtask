package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error response.
func Error(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

// JSON response.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
	}
}
