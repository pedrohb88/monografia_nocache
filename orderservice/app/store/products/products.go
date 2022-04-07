package products

import (
	"database/sql"
	"fmt"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryProductsBase = `
	SELECT 
		p.id AS ID,
		p.name AS Name,
		p.price AS Price
	FROM products p
	%s`

	all = fmt.Sprintf(queryProductsBase, "")

	byID = fmt.Sprintf(queryProductsBase, `WHERE p.id = ?`)

	execInsert = `
	INSERT INTO products(name, price)
	VALUES (?, ?)
	`

	execDelete = `DELETE FROM products WHERE id = ?`
)

type Products interface {
	GetAll() ([]*model.Product, error)
	GetByID(productID int) (*model.Product, error)
	Create(product *model.Product) error
	Delete(productID int) error
}

type products struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Products {
	return &products{db: db}
}

func (p *products) GetAll() ([]*model.Product, error) {
	var products []*model.Product

	_, err := p.db.Select(&products, all)
	return products, err
}

func (p *products) GetByID(productID int) (*model.Product, error) {
	var product model.Product
	err := p.db.SelectOne(&product, byID, productID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *products) Create(product *model.Product) error {

	res, err := p.db.Exec(execInsert,
		product.Name,
		product.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	product.ID = int(lastID)
	return nil
}

func (p *products) Delete(productID int) error {
	res, err := p.db.Exec(execDelete, productID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		return nil
	}
	return sql.ErrNoRows
}
