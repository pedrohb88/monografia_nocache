package service

import (
	"monografia/model"
	"monografia/store/products"
	"monografia/transport/entity"
)

type productsService struct {
	productsStore products.Products
}

func (p *productsService) GetAll() ([]*entity.Product, error) {

	productsModels, err := p.productsStore.GetAll()
	if err != nil {
		return nil, err
	}

	return entity.NewProducts(productsModels), nil
}

func (p *productsService) GetByID(productID int) (*entity.Product, error) {

	productModel, err := p.productsStore.GetByID(productID)
	if err != nil {
		return nil, err
	}

	return entity.NewBasicProduct(productModel), nil
}

func (p *productsService) Create(product *model.Product) error {
	return p.productsStore.Create(product)
}

func (p *productsService) Delete(productID int) error {
	return p.productsStore.Delete(productID)
}
