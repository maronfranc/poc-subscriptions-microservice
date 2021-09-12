package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Listen(r *mux.Router, port uint) {
	r.Use(mux.CORSMethodMiddleware(r))
	http.Handle("/", r)
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
