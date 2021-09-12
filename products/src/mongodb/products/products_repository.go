package products

import (
	"github.com/maronfranc/subscription-system-products/src/models"
	"github.com/maronfranc/subscription-system-products/src/mongodb"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "products"

// FindManyAndCount find products and count
func FindManyAndCount(page int64, perPage int64, filter interface{}) ([]*models.ProductModel, int64, error) {
	var pp []*models.ProductModel
	skip := (page - 1) * perPage
	o := options.Find().SetLimit(perPage).SetSkip(skip)
	err := mongodb.FindAll(COLLECTION_NAME, &pp, filter, o)
	if err != nil {
		return pp, 0, err
	}

	count, err := mongodb.CountDocuments(COLLECTION_NAME, options.Count())
	return pp, count, err
}

// FindOne
func FindOne(filter interface{}) (models.ProductModel, error) {
	var p = models.ProductModel{}
	err := mongodb.FindOne(COLLECTION_NAME, filter, &p)
	return p, err
}

// Insert new product
func Insert(name string, price uint) error {
	p := models.ProductModel{
		Name:  name,
		Price: price,
	}
	_, err := mongodb.Insert(COLLECTION_NAME, p)
	return err
}

// Delete
func Delete(filter interface{}) (bool, error) {
	deleted, err := mongodb.Delete(COLLECTION_NAME, filter)
	return deleted, err
}
