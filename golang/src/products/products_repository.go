package products

import (
	"github.com/maronfranc/subscription-system-products/src/models"
)

func GetMany(page uint, per_page uint) []models.PaginatedProductModel {
	ps := []models.ProductModel{
		{
			Id:    "123123123",
			Name:  "Name",
			Price: 100,
		},
		{
			Id:    "222222222",
			Name:  "Name 2",
			Price: 250,
		},
	}
	count := uint(25)
	many_products := []models.PaginatedProductModel{
		{
			Data:       ps,
			Page:       page,
			TotalPages: count / per_page,
		},
	}
	return many_products
}

func GetById(id string) models.ProductModel {
	p := models.ProductModel{
		Id:    id,
		Name:  "Name",
		Price: 100,
	}
	return p
}

func CreateProduct(name string, price uint) models.ProductModel {
	p := models.ProductModel{
		Id:    "987654321",
		Name:  name,
		Price: price,
	}
	return p
}
