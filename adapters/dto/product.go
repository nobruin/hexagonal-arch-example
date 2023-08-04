package dto

import "github.com/nobruin/hexagonal-arch-example/app"

type ProductRequest struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProductRequest() *ProductRequest {
	return &ProductRequest{}
}

func (r *ProductRequest) toEntity(product *app.Product) (*app.Product, error) {
	if r.ID != "" {
		product.ID = r.ID
	}
	product.Name = r.Name
	product.Price = r.Price
	product.Status = r.Status

	_, err := product.IsValid()
	if err != nil {
		return &app.Product{}, err
	}
	return product, nil
}
