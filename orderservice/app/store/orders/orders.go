package orders

import (
	"fmt"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryOrdersBase = `
	SELECT 
		o.id AS ID, 
		o.user_id AS UserID, 
		o.items_quantity AS ItemsQuantity, 
		o.price AS Price,
		i.id AS ItemID,
		i.quantity AS ItemQuantity,
		i.price AS ItemPrice,
		p.id AS ProductID,
		p.name AS ProductName,
		p.price AS ProductPrice
	FROM orders o
	INNER JOIN items i
		ON i.order_id = o.id
	INNER JOIN products p
		ON i.product_id = p.id
	%s
	`

	byUserID = fmt.Sprintf(queryOrdersBase, `WHERE o.user_id = ?`)

	byID = fmt.Sprintf(queryOrdersBase, `WHERE o.id = ?`)

	execInsertOrder = `
	INSERT INTO orders(user_id, items_quantity, price)
	VALUES (?, ?, ?)
	`
)

type Orders interface {
	GetByUserID(userID int) ([]model.Order, error)
	GetByID(orderID int) ([]model.Order, error)
	Create(order *model.Order) error
}

type orders struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Orders {
	return &orders{db: db}
}

func (o *orders) GetByUserID(userID int) ([]model.Order, error) {
	var orders []model.Order

	_, err := o.db.Select(&orders, byUserID, userID)
	return orders, err
}

func (o *orders) GetByID(orderID int) ([]model.Order, error) {
	var orders []model.Order

	_, err := o.db.Select(&orders, byID, orderID)
	return orders, err
}

func (o *orders) Create(order *model.Order) error {

	res, err := o.db.Exec(execInsertOrder,
		order.UserID,
		order.ItemsQuantity,
		order.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	order.ID = int(lastID)
	return nil
}
