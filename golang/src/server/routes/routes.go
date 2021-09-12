package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	Products(r)
	return r
}
