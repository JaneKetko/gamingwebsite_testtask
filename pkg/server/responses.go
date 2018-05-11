package server

import (
	"encoding/json"
	"net/http"
)

//Error response
func Error(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

//JSON response
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
