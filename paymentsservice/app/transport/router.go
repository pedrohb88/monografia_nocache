package transport

import (
	"monografia/service"

	"github.com/go-chi/chi/v5"
)

func NewRouter(srv service.Service) *chi.Mux {
	r := chi.NewRouter()

	payments := &Payments{service: &srv}
	invoices := &Invoices{service: &srv}

	// Payments
	r.Get("/payments/{paymentID:[1-9][0-9]*}", payments.GetByID)
	r.Post("/payments", payments.Create)

	// Invoices
	r.Get("/invoices/{invoiceID:[1-9][0-9]*}", invoices.GetByID)
	r.Post("/invoices", invoices.Create)

	return r
}
