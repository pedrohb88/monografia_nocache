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

type Payments struct {
	service *service.Service
}

func (o *Payments) GetByID(w http.ResponseWriter, r *http.Request) {
	paymentID := util.ParamAsInt(r, "paymentID")

	payment, err := o.service.Payments.GetByID(paymentID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, payment)
}

func (o *Payments) Create(w http.ResponseWriter, r *http.Request) {
	var paymentModel model.Payment
	err := util.DecodeJSON(r, &paymentModel)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = o.service.Payments.Create(&paymentModel)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, entity.NewBasicPayment(paymentModel))
}
