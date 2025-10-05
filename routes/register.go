package routes

import (
	"go-api-native/controllers/authcontroller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	router := r.PathPrefix("/register").Subrouter()

	router.HandleFunc("", authcontroller.Register).Methods("POST")

}
