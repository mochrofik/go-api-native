package routes

import (
	"go-api-native/controllers/authcontroller"
	"go-api-native/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("/me", authcontroller.Me).Methods("GET")
}
