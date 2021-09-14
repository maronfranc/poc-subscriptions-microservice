package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maronfranc/subscription-system-products/src/models"
	"github.com/maronfranc/subscription-system-products/src/mongodb/products"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert creates a new products
func Insert(w http.ResponseWriter, r *http.Request) {
	var pdt models.ProductModel
	err := json.NewDecoder(r.Body).Decode(&pdt)
	if err != nil {
		http.Error(w, "Invalid product", http.StatusBadRequest)
		return
	}

	p := products.Insert(pdt.Name, pdt.Price)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// GetPaginated return paginated products
func GetPaginated(w http.ResponseWriter, r *http.Request) {
	pg, err := queryString(r, "page")
	if err != nil {
		http.Error(w, "Invalid query parameter type {page}", http.StatusBadRequest)
		return
	}

	perPage, err := queryString(r, "per_page")
	if err != nil {
		http.Error(w, "Invalid query parameter type {page}", http.StatusBadRequest)
		return
	}

	pp, count, err := products.FindManyAndCount(pg, perPage, bson.M{})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	paginated := models.PaginatedProductModel{
		Data:       pp,
		Page:       pg,
		TotalPages: (count / perPage),
		TotalItems: count,
		Count:      int64(len(pp)),
	}
	if err := json.NewEncoder(w).Encode(paginated); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// GetById
func GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objId}}
	p, err := products.FindOne(filter)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

// queryStringInt gets query params from url and convert to int
func queryStringInt(r *http.Request, queryName string) (int64, error) {
	str := r.FormValue(queryName)
	if str != "" {
		pg, err := strconv.ParseInt(str, 10, 32)
		return pg, err
	}
	return 0, nil
}

func queryString(r *http.Request, queryName string) (int64, error) {
	pg, err := queryStringInt(r, queryName)
	if pg == 0 {
		switch queryName {
		case "page":
			return DEFAULT_PAGE, err
		case "per_page":
			return DEFAULT_PER_PAGE, err
		}
	}
	return pg, err
}
