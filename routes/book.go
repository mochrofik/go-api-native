package routes

import (
	"go-api-native/controllers/bookcontroller"
	"go-api-native/middleware"

	"github.com/gorilla/mux"
)

func BooksRoutes(r *mux.Router) {

	router := r.PathPrefix("/books").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("", bookcontroller.Index).Methods("GET")
	router.HandleFunc("", bookcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", bookcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", bookcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", bookcontroller.Delete).Methods("DELETE")
}
