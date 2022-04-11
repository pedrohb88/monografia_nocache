package service

import (
	"monografia/store/invoices"
	"monografia/store/payments"
)

type Service struct {
	Payments paymentsService
	Invoices invoicesService
}

func New(
	paymentsStore payments.Payments,
	invoicesStore invoices.Invoices,
) Service {

	paymentsService := paymentsService{
		paymentsStore: paymentsStore,
		invoicesStore: invoicesStore,
	}

	invoicesService := invoicesService{
		invoicesStore: invoicesStore,
	}

	return Service{
		Payments: paymentsService,
		Invoices: invoicesService,
	}
}
