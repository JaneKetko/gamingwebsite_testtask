package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO add unit tests to Error and JSON.

// Error is error response.
func Error(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

// JSON is success response with http code and json(payload) as message.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
	}
}
