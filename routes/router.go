package routes

import (
	"api/config"
	UserController "api/controllers/user"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.Use(config.Cors)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the blank space")
	})

	api := r.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/user").Subrouter()
	user.HandleFunc("/listUser", UserController.UsersList).Methods("GET")

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
