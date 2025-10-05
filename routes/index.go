package routes

import "github.com/gorilla/mux"

func RouteIndex(r *mux.Router) {

	api := r.PathPrefix("/api").Subrouter()

	AuthRoutes(api)
	RegisterRoutes(api)
	AuthorRoutes(api)
	BooksRoutes(api)
}
