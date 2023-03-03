package user

import (
	"api/config"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetListUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": err,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "success",
		"payload": users,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": err,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
			return
		}
		return
	}
}
