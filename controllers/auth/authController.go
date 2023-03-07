package auth

import (
	AuthServices "api/services/auth"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	AuthServices.Login(w, r)
}
