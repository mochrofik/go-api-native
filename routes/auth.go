package routes

import (
	"go-api-native/controllers/authcontroller"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	route := r.PathPrefix("/auth").Subrouter()

	route.HandleFunc("/login", authcontroller.Login).Methods("POST")
}
