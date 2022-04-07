package items

import (
	"database/sql"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	execInsertItem = `
	INSERT INTO items(order_id, product_id, quantity, price) 
	VALUES(?, ?, ?, ?)
	`

	execDeleteItem = `DELETE FROM items WHERE id = ?`
)

type Items interface {
	Create(item *model.Item) error
	Delete(itemID int) error
}

type items struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Items {
	return &items{db: db}
}

func (i *items) Create(item *model.Item) error {
	res, err := i.db.Exec(execInsertItem,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.Price,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	item.ID = int(lastID)
	return nil
}

func (i *items) Delete(itemID int) error {
	res, err := i.db.Exec(execDeleteItem, itemID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		return nil
	}
	return sql.ErrNoRows
}
