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

var secret = os.Getenv("JWT_SECRET")

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		if tokenString == "" {
			common.RespondWithError(w, http.StatusUnauthorized, "Token Invalid")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})
		if err != nil {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized request")
			return
		}

		if !token.Valid {
			common.RespondWithError(w, http.StatusUnauthorized, "Token not found")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().Unix() > exp {
				common.RespondWithError(w, http.StatusUnauthorized, "Token expired")
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
