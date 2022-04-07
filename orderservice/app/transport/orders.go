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

type Orders struct {
	service *service.Service
}

func (o *Orders) GetByUserID(w http.ResponseWriter, r *http.Request) {

	userID := util.ParamAsInt(r, "userID")

	orders, err := o.service.Orders.GetByUserID(userID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, orders)
}

func (o *Orders) GetByID(w http.ResponseWriter, r *http.Request) {
	orderID := util.ParamAsInt(r, "orderID")

	order, err := o.service.Orders.GetByID(orderID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, order)
}

func (o *Orders) Create(w http.ResponseWriter, r *http.Request) {
	var orderModel model.Order
	err := util.DecodeJSON(r, &orderModel)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = o.service.Orders.Create(&orderModel)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, entity.NewBasicOrder(orderModel))
}

func (o *Orders) AddItem(w http.ResponseWriter, r *http.Request) {

	orderID := util.ParamAsInt(r, "orderID")

	var item model.Item
	err := util.DecodeJSON(r, &item)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	item.OrderID = orderID

	err = o.service.Orders.AddItem(&item)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	order, err := o.service.Orders.GetByID(orderID)
	if err != nil {
		response.Empty(w, http.StatusOK)
		return
	}

	response.JSON(w, order)
}

func (o *Orders) RemoveItem(w http.ResponseWriter, r *http.Request) {
	orderID := util.ParamAsInt(r, "orderID")
	itemID := util.ParamAsInt(r, "itemID")

	err := o.service.Orders.RemoveItem(itemID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	order, err := o.service.Orders.GetByID(orderID)
	if err != nil {
		response.Empty(w, http.StatusOK)
		return
	}

	response.JSON(w, order)
}

func (o *Orders) Pay(w http.ResponseWriter, r *http.Request) {
	// Esse cara vai chamar o outro microsservice, recebendo o ID do payment como resposta pra atualizar a order
	return
}
