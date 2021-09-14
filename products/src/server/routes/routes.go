package routes

import (
	"github.com/gorilla/mux"
)

// LoadRoutes
func LoadRoutes(r *mux.Router) {
	Products(r)
}
