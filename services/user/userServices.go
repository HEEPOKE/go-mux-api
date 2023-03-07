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
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, "Success", user)
}
