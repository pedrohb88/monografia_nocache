package service

import (
	"monografia/errors"
	"monografia/model"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/products"
)

type ordersService struct {
	ordersStore   orders.Orders
	itemsStore    items.Items
	productsStore products.Products
}

func (o *ordersService) GetByUserID(userID int) ([]*model.Order, error) {
	return o.ordersStore.GetByUserID(userID)
}

func (o *ordersService) GetByID(orderID int) (*model.Order, error) {
	return o.ordersStore.GetByID(orderID)
}

func (o *ordersService) Create(order *model.Order) error {
	return o.ordersStore.Create(order)
}

func (o *ordersService) GetItems(orderID int) ([]*model.Item, error) {
	return o.itemsStore.GetByOrderID(orderID)
}

func (o *ordersService) GetItemByID(itemID int) (*model.Item, error) {
	return o.itemsStore.GetByID(itemID)
}

func (o *ordersService) AddItem(item *model.Item) error {

	_, err := o.productsStore.GetByID(item.ProductID)
	if errors.IsNotFound(err) {
		return errors.ErrProductNotFound
	}
	if err != nil {
		return err
	}

	return o.itemsStore.Create(item)
}

func (o *ordersService) RemoveItem(itemID int) error {
	return o.itemsStore.Delete(itemID)
}
