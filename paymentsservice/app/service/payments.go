package service

import (
	"database/sql"
	"monografia/model"
	"monografia/store/invoices"
	"monografia/store/payments"
	"monografia/transport/entity"
)

type paymentsService struct {
	paymentsStore payments.Payments
	invoicesStore invoices.Invoices
}

func (o *paymentsService) GetByID(paymentID int) (*entity.Payment, error) {
	paymentModels, err := o.paymentsStore.GetByID(paymentID)
	if err != nil {
		return nil, err
	}
	if len(paymentModels) == 0 {
		return nil, sql.ErrNoRows
	}

	return entity.NewPayments(paymentModels)[0], nil
}

func (o *paymentsService) Create(payment *model.Payment) error {
	return o.paymentsStore.Create(payment)
}
