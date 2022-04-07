package orders

import (
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryOrdersByUserID = `SELECT * FROM orders WHERE user_id = ?`

	queryOrderByID = `SELECT * FROM orders WHERE id = ?`

	execInsertOrder = `
	INSERT INTO orders(user_id, items_quantity, price)
	VALUES (?, ?, ?)
	`
)

type Orders interface {
	GetByUserID(userID int) ([]*model.Order, error)
	GetByID(orderID int) (*model.Order, error)
	Create(order *model.Order) error
}

type orders struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Orders {
	return &orders{db: db}
}

func (o *orders) GetByUserID(userID int) ([]*model.Order, error) {
	var orders []*model.Order

	_, err := o.db.Select(&orders, queryOrdersByUserID, userID)
	return orders, err
}

func (o *orders) GetByID(orderID int) (*model.Order, error) {
	var order *model.Order

	err := o.db.SelectOne(&order, queryOrderByID, orderID)
	return order, err
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
