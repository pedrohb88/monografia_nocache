package entity

import "monografia/model"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewBasicProduct(p *model.Product) *Product {
	return &Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

func NewProducts(models []*model.Product) []*Product {
	products := make([]*Product, len(models))
	for i, m := range models {
		products[i] = NewBasicProduct(m)
	}
	return products
}
