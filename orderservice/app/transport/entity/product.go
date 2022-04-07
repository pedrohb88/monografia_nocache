package entity

import "monografia/model"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (e *Entity) NewProduct(p *model.Product) *Product {
	return &Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

func (e *Entity) NewProductByID(productID int) (*Product, error) {

	product, err := e.service.Products.GetByID(productID)
	if err != nil {
		return nil, err
	}

	return e.NewProduct(product), nil
}
