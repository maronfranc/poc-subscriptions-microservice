package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maronfranc/subscription-system-products/src/products"
)

func HandleGetObject(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandleGetObject)
	products.Routes(r)
	http.Handle("/", r)
	port := "8080"
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
