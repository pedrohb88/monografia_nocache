package payments

import (
	"fmt"
	"monografia/model"

	"github.com/go-gorp/gorp"
)

var (
	queryPaymentsBase = `
	SELECT 
		p.id AS ID, 
		p.amount AS Amount,
		p.invoice_id AS InvoiceID, 
		i.code AS InvoiceCode, 
		i.link AS InvoiceLink
	FROM payments p
	INNER JOIN invoices i
		ON p.invoice_id = i.id
	%s
	`

	byID = fmt.Sprintf(queryPaymentsBase, `WHERE p.id = ?`)

	execInsertPayment = `
	INSERT INTO payments(amount)
	VALUES (?)
	`
)

type Payments interface {
	GetByID(paymentID int) ([]model.Payment, error)
	Create(payment *model.Payment) error
}

type payments struct {
	db *gorp.DbMap
}

func New(db *gorp.DbMap) Payments {
	return &payments{db: db}
}

func (o *payments) GetByID(paymentID int) ([]model.Payment, error) {
	var payments []model.Payment

	_, err := o.db.Select(&payments, byID, paymentID)
	return payments, err
}

func (o *payments) Create(payment *model.Payment) error {

	res, err := o.db.Exec(execInsertPayment,
		payment.Amount,
	)
	if err != nil {
		return err
	}

	lastID, _ := res.LastInsertId()
	payment.ID = int(lastID)
	return nil
}
