package model

type Payment struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	InvoiceID   int     `json:"invoice_id"`
	InvoiceCode string  `json:"invoice_code"`
	InvoiceLink string  `json:"invoice_link"`
}
