package invoices

import (
	"fmt"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryInvoicesBase = `
	SELECT 
		i.id AS ID,
		i.code AS Code,
		i.link AS Link
	FROM invoices i
	%s`

	byID = fmt.Sprintf(queryInvoicesBase, `WHERE i.id = ?`)

	execInsert = `
	INSERT INTO invoices(code, link)
	VALUES (?, ?)
	`
)

type Invoices interface {
	GetByID(invoiceID int) (*model.Invoice, error)
	Create(invoice *model.Invoice) error
}

type invoices struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Invoices {
	return &invoices{db: db}
}

func (p *invoices) GetByID(invoiceID int) (*model.Invoice, error) {
	var invoice model.Invoice
	err := p.db.SelectOne(&invoice, byID, invoiceID)
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (p *invoices) Create(invoice *model.Invoice) error {

	res, err := p.db.Exec(execInsert,
		invoice.Code,
		invoice.Link,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	invoice.ID = int(lastID)
	return nil
}
