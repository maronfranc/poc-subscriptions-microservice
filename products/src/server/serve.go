package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maronfranc/subscription-system-products/src/server/routes"
)

// Listen load routes and run server
func Listen(port uint) {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	r.Use(mux.CORSMethodMiddleware(r))
	routes.LoadRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
