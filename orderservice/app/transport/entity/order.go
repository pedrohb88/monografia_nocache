package entity

import "monografia/model"

type Order struct {
	ID            int         `json:"id"`
	UserID        int         `json:"user_id"`
	ItemsQuantity int         `json:"items_quantity"`
	Price         float64     `json:"price"`
	Items         []*Item     `json:"items"`
	Payment       interface{} `json:"payment"`
}

func (e *Entity) NewOrder(m model.Order) Order {
	return Order{
		ID:            m.ID,
		UserID:        m.UserID,
		ItemsQuantity: m.ItemsQuantity,
		Price:         m.Price,
	}
}

func (e *Entity) NewOrderByID(orderID int) (*Order, error) {

	order, err := e.service.Orders.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	itemsModel, err := e.service.Orders.GetItems(orderID)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, len(itemsModel))
	for i, item := range itemsModel {
		it, err := e.NewItemByID(item.ID)
		if err != nil {
			return nil, err
		}
		items[i] = it
	}

	// TODO : Buscar PAYMENT
	o := e.NewOrder(*order)
	o.Items = items

	return &o, nil
}
