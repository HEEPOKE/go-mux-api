package user

import (
	"api/common"
	"api/config"
	"api/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetListUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "success", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var users models.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	users.Password = string(hashedPassword)

	config.DB.Create(&users)
	common.RespondWithJSON(w, http.StatusCreated, "Success", users)
}
