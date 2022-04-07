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

type Products struct {
	service *service.Service
}

func (p *Products) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.service.Products.GetAll()
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, products)
}

func (p *Products) GetByID(w http.ResponseWriter, r *http.Request) {
	productID := util.ParamAsInt(r, "productID")

	product, err := p.service.Products.GetByID(productID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, product)
}

func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	var productModel model.Product
	err := util.DecodeJSON(r, &productModel)
	if err != nil {
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	err = p.service.Products.Create(&productModel)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, entity.NewBasicProduct(&productModel))
}

func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	productID := util.ParamAsInt(r, "productID")

	err := p.service.Products.Delete(productID)
	if errors.IsNotFound(err) {
		response.Empty(w, http.StatusNotFound)
		return
	}
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}
}
