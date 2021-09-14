package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maronfranc/subscription-system-products/src/server/controllers"
)

// Products routes
func Products(r *mux.Router) {
	route := r.PathPrefix("/products").Subrouter()
	route.Path("").Methods(http.MethodPost).HandlerFunc(controllers.Insert)
	route.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controllers.GetById)
	route.Path("").Methods(http.MethodGet).HandlerFunc(controllers.GetPaginated)
}
