package models

type ProductModel struct {
	Id   interface{} `json:"_id" bson:"_id"`
	Name string      `json:"name" bson:"name"`
	// Price in cents
	Price uint `json:"price" bson:"price"`
}

type PaginatedProductModel struct {
	Data       []*ProductModel `json:"data" bson:"data"`
	TotalItems int64           `json:"total_items" bson:"total_items"`
	TotalPages int64           `json:"total_pages" bson:"total_pages"`
	Count      int64           `json:"count" bson:"count"`
	Page       int64           `json:"page" bson:"page"`
}
