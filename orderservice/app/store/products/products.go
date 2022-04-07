package products

import (
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryAllProducts = `SELECT * FROM products`

	queryByID = `SELECT * FROM products WHERE id = ?`
)

type Products interface {
	GetAll() ([]*model.Product, error)
	GetByID(productID int) (*model.Product, error)
}

type products struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Products {
	return &products{db: db}
}

func (o *products) GetAll() ([]*model.Product, error) {
	var products []*model.Product

	_, err := o.db.Select(&products, queryAllProducts)
	return products, err
}

func (o *products) GetByID(productID int) (*model.Product, error) {
	var product model.Product
	err := o.db.SelectOne(&product, queryByID, productID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
