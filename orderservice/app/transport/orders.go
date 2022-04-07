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
	entity  *entity.Entity
}

func (o *Orders) GetByUserID(w http.ResponseWriter, r *http.Request) {

	userID := util.ParamAsInt(r, "userID")

	ordersModel, err := o.service.Orders.GetByUserID(userID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	orders := make([]*entity.Order, len(ordersModel))
	for i, order := range ordersModel {

		o, err := o.entity.NewOrderByID(order.ID)
		if err != nil {
			response.Err(w, http.StatusInternalServerError, err)
			return
		}

		orders[i] = o
	}

	response.JSON(w, orders)
}

func (o *Orders) GetByID(w http.ResponseWriter, r *http.Request) {
	orderID := util.ParamAsInt(r, "orderID")

	order, err := o.entity.NewOrderByID(orderID)
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

	var order model.Order
	err := util.DecodeJSON(r, &order)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = o.service.Orders.Create(&order)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, o.entity.NewOrder(order))
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

	order, err := o.entity.NewOrderByID(item.OrderID)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, order)
}

func (o *Orders) RemoveItem(w http.ResponseWriter, r *http.Request) {
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
}

func (o *Orders) Pay(w http.ResponseWriter, r *http.Request) {
	// Esse cara vai chamar o outro microsservice, recebendo o ID do payment como resposta pra atualizar a order
	return
}
