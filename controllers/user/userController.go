package user

import (
	UserServices "api/services/user"
	"net/http"
)

func UsersList(w http.ResponseWriter, r *http.Request) {
	UserServices.GetListUser(w, r)
}
