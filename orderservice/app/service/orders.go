package service

import (
	"database/sql"
	"monografia/errors"
	"monografia/model"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/products"
	"monografia/transport/entity"
)

type ordersService struct {
	ordersStore   orders.Orders
	itemsStore    items.Items
	productsStore products.Products
}

func (o *ordersService) GetByUserID(userID int) ([]*entity.Order, error) {

	ordersModels, err := o.ordersStore.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return entity.NewOrders(ordersModels), nil
}

func (o *ordersService) GetByID(orderID int) (*entity.Order, error) {
	orderModels, err := o.ordersStore.GetByID(orderID)
	if len(orderModels) == 0 {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}

	return entity.NewOrders(orderModels)[0], nil
}

func (o *ordersService) Create(order *model.Order) error {
	return o.ordersStore.Create(order)
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
