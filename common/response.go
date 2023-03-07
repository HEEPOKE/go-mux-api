package common

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, message, nil)
}

func RespondWithJSON(w http.ResponseWriter, code int, message string, payload ...interface{}) {
	var response map[string]interface{}
	if len(payload) > 0 {
		response = map[string]interface{}{"message": message, "status": code, "payload": payload[0]}
	} else {
		response = map[string]interface{}{"message": message, "status": code}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
