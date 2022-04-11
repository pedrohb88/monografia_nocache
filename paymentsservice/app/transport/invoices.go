package transport

import (
	"monografia/errors"
	"monografia/model"
	"monografia/service"
	"monografia/transport/entity"
	"monografia/transport/util"
	"monografia/transport/util/response"
	"net/http"
)

type Invoices struct {
	service *service.Service
}

func (p *Invoices) GetByID(w http.ResponseWriter, r *http.Request) {
	invoiceID := util.ParamAsInt(r, "invoiceID")

	invoice, err := p.service.Invoices.GetByID(invoiceID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, invoice)
}

func (p *Invoices) Create(w http.ResponseWriter, r *http.Request) {
	var invoiceModel model.Invoice
	err := util.DecodeJSON(r, &invoiceModel)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = p.service.Invoices.Create(&invoiceModel)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, entity.NewBasicInvoice(&invoiceModel))
}
