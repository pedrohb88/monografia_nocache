package entity

import (
	"monografia/model"
)

type Item struct {
	ID       int      `json:"id"`
	OrderID  int      `json:"order_id"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
}

func (e *Entity) NewItem(i model.Item) *Item {
	return &Item{
		ID:       i.ID,
		OrderID:  i.OrderID,
		Quantity: i.Quantity,
		Price:    i.Price,
	}
}

func (e *Entity) NewItemByID(itemID int) (*Item, error) {

	itemModel, err := e.service.Orders.GetItemByID(itemID)
	if err != nil {
		return nil, err
	}

	product, _ := e.NewProductByID(itemModel.ProductID)

	item := e.NewItem(*itemModel)
	item.Product = product

	return item, nil
}
