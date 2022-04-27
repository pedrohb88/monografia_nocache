package service

import (
	"database/sql"
	"monografia/errors"
	"monografia/model"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/payments"
	"monografia/store/products"
	"monografia/transport/entity"
)

type ordersService struct {
	ordersStore   orders.Orders
	itemsStore    items.Items
	productsStore products.Products
	paymentsStore payments.Payments
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

	e := entity.NewOrders(orderModels)[0]

	var payment *model.Payment 
	if orderModels[0].PaymentID != nil {
		payment, err = o.paymentsStore.GetByID(*orderModels[0].PaymentID)
		if err != nil {
			return nil, err
		}

		e.Payment = &entity.Payment{
			ID:     payment.ID,
			Amount: payment.Amount,
			Invoice: &entity.Invoice{
				ID:   payment.Invoice.ID,
				Code: payment.Invoice.Code,
				Link: payment.Invoice.Link,
			},
		}
	}

	return e, nil
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

func (o *ordersService) Pay(orderID int, amount float64) error {
	paymentID, err := o.paymentsStore.Create(amount)
	if err != nil {
		return err
	}

	return o.ordersStore.UpdatePaymentID(orderID, paymentID)
}
