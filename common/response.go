package common

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, message, nil)
}

func RespondWithJSON(w http.ResponseWriter, code int, message string, payload interface{}) {
	response := map[string]interface{}{"message": message, "status": code, "payload": payload}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
