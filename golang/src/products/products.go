package products

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	productsR := r.PathPrefix("/products").Subrouter()
	productsR.Path("").Methods(http.MethodPost).HandlerFunc(createProduct)
	productsR.Path("").Methods(http.MethodGet).HandlerFunc(getProducts)
	productsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getProductById)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Create products</h1>"))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Get products</h1>"))
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Get product: " + id + "</h1>"))
}
