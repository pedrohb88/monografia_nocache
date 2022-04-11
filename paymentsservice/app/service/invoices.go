package service

import (
	"monografia/model"
	"monografia/store/invoices"
	"monografia/transport/entity"
)

type invoicesService struct {
	invoicesStore invoices.Invoices
}

func (p *invoicesService) GetByID(invoiceID int) (*entity.Invoice, error) {

	invoiceModel, err := p.invoicesStore.GetByID(invoiceID)
	if err != nil {
		return nil, err
	}

	return entity.NewBasicInvoice(invoiceModel), nil
}

func (p *invoicesService) Create(invoice *model.Invoice) error {
	return p.invoicesStore.Create(invoice)
}
