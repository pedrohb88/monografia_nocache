package entity

import "monografia/model"

type Order struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	ItemsQuantity int     `json:"items_quantity"`
	Price         float64 `json:"price"`
	Items         []*Item `json:"items"`
}

func NewBasicOrder(model model.Order) *Order {
	return &Order{
		ID:            model.ID,
		UserID:        model.UserID,
		ItemsQuantity: model.ItemsQuantity,
		Price:         model.Price,
	}
}

func NewOrders(models []model.Order) []*Order {

	var orders []*Order

	var currentOrder *Order
	for _, model := range models {

		if currentOrder == nil || model.ID != currentOrder.ID {
			currentOrder = NewBasicOrder(model)
			orders = append(orders, currentOrder)
		}

		if model.ItemID == nil {
			continue
		}

		item := &Item{
			ID:       *model.ItemID,
			OrderID:  model.ID,
			Quantity: *model.ItemQuantity,
			Price:    *model.ItemPrice,
		}

		product := &Product{
			ID:    *model.ProductID,
			Name:  *model.ProductName,
			Price: *model.ProductPrice,
		}

		item.Product = product

		currentOrder.Items = append(currentOrder.Items, item)
	}

	return orders
}
