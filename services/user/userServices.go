package user

import (
	"api/common"
	"api/config"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetListUser(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	err := config.DB.Find(&user).Error
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "success", user)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var oldUser models.User
	if err := config.DB.Where("id = ?", param).First(&oldUser).Error; err != nil {
		common.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := config.DB.Model(&oldUser).Updates(user).Error; err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "Success")
}
