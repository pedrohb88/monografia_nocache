package entity

import "monografia/model"

type Payment struct {
	ID      int      `json:"id"`
	Amount  float64  `json:"amount"`
	Invoice *Invoice `json:"invoice"`
}

func NewBasicPayment(model model.Payment) *Payment {
	return &Payment{
		ID:     model.ID,
		Amount: model.Amount,
	}
}

func NewPayments(models []model.Payment) []*Payment {

	var payments []*Payment

	for _, model := range models {

		payment := NewBasicPayment(model)
		payment.Invoice = &Invoice{
			ID:   model.InvoiceID,
			Code: model.InvoiceCode,
			Link: model.InvoiceLink,
		}

		payments = append(payments, payment)
	}

	return payments
}
