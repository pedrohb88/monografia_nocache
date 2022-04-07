package service

import (
	"monografia/model"
	"monografia/store/products"
)

type productsService struct {
	productsStore products.Products
}

func (o *productsService) GetByID(productID int) (*model.Product, error) {
	return o.productsStore.GetByID(productID)
}
