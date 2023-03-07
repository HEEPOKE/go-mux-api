package auth

import (
	"api/common"
	"api/config"
	"api/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var auth models.Auth
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user models.User
	if err = config.DB.Where("username = ?", auth.Username).First(&user).Error; err != nil {
		common.RespondWithError(w, http.StatusUnauthorized, "Username is incorrect")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		common.RespondWithError(w, http.StatusUnauthorized, "password is incorrect")
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		common.RespondWithError(w, http.StatusInternalServerError, "JWT secret is not set")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "Success", map[string]string{
		"token": tokenString,
	})
}
