package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Listen(r *mux.Router, port uint) {
	http.Handle("/", r)
	fmt.Println("Starting up on ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
