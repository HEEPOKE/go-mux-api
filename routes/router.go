package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the home page!")
	})

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of products")
	})

	r.HandleFunc("/products/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		productID := vars["id"]
		fmt.Fprintf(w, "Showing product with ID %v", productID)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
