package routes

import (
	"api/config"
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

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
	fmt.Println("http:\\localhost:" + os.Getenv("PORT"))
}
