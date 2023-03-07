package user

import (
	UserServices "api/services/user"
	"net/http"
)

func UsersList(w http.ResponseWriter, r *http.Request) {
	UserServices.GetListUser(w, r)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	UserServices.CreateUser(w, r)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	UserServices.UpdateUser(w, r)
}
