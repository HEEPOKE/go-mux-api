package middleware

import (
	"api/common"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		tokenString := strings.Split(header, " ")
		fmt.Println(tokenString[1])
		if len(tokenString) != 2 {
			common.RespondWithError(w, http.StatusUnauthorized, "Token Invalid")
			return
		}
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})

		if err != nil || !token.Valid {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		exp := int64(claims["exp"].(float64))

		if time.Now().Unix() > exp {
			common.RespondWithError(w, http.StatusUnauthorized, "Token expired")
			return
		}

		next.ServeHTTP(w, r)
	})
}
