package server

import (
	"encoding/json"
	"net/http"
)

// TODO points in the end of comments?
//Error response
func Error(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

//JSON response
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	// TODO why don't you process error?
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// TODO why don't you process error? Show this error in stderr
	w.Write(response)
}
