package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PER_PAGE = 10

func Routes(r *mux.Router) {
	productsR := r.PathPrefix("/products").Subrouter()
	productsR.Path("").Methods(http.MethodPost).HandlerFunc(createProduct)
	productsR.Path("").Queries("page", "{page:[0-9]+}").Methods(http.MethodGet).HandlerFunc(getProducts)
	productsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getProductById)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	product := mux.Vars(r)
	price, err := strconv.ParseUint(product["price"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
	}
	p := CreateProduct(product["name"], uint(price))
	if err := json.NewEncoder(w).Encode(p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	page_string := r.FormValue("page")
	page, err := strconv.ParseUint(page_string, 10, 32)
	if err != nil && page_string != "" {
		http.Error(w, "Invalid query parameter type {page}", http.StatusBadRequest)
	}
	ps := GetMany(uint(page), PER_PAGE)
	if err := json.NewEncoder(w).Encode(ps); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	p := GetById(id)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}
