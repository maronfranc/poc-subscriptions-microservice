package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maronfranc/subscription-system-products/src/products"
)

func main() {
	r := mux.NewRouter()
	products.Routes(r)
	http.Handle("/", r)
	port := "8080"
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
