package routes

import (
	"api/config"
	AuthController "api/controllers/auth"
	UserController "api/controllers/user"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(port string) {
	r := mux.NewRouter()
	r.Use(config.Cors)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the blank space")
	})

	api := r.PathPrefix("/api").Subrouter()
	auth := api.PathPrefix("/auth").Subrouter()
	user := api.PathPrefix("/user").Subrouter()

	auth.HandleFunc("/login", AuthController.Login).Methods("POST")

	user.HandleFunc("/", UserController.UsersList).Methods("GET")
	user.HandleFunc("/add", UserController.UserCreate).Methods("POST")
	user.HandleFunc("/update/{id}", UserController.UserUpdate).Methods("PUT")
	user.HandleFunc("/delete/{id}", UserController.UserDelete).Methods("DELETE")

	http.ListenAndServe(":"+port, r)
}
