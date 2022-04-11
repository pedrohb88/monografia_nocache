package entity

import "monografia/model"

type Invoice struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Link string `json:"link"`
}

func NewBasicInvoice(i *model.Invoice) *Invoice {
	return &Invoice{
		ID:   i.ID,
		Code: i.Code,
		Link: i.Link,
	}
}

func NewInvoices(models []*model.Invoice) []*Invoice {
	invoices := make([]*Invoice, len(models))
	for i, m := range models {
		invoices[i] = NewBasicInvoice(m)
	}
	return invoices
}
