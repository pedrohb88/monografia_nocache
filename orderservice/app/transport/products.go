package transport

import (
	"monografia/service"
	"monografia/transport/entity"
	"net/http"
)

type Products struct {
	service *service.Service
	entity  *entity.Entity
}

func (p *Products) GetByID(w http.ResponseWriter, req *http.Request) {
	return
}

func (p *Products) Create(w http.ResponseWriter, req *http.Request) {
	return
}
