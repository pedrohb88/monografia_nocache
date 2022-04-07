package transport

import (
	"monografia/service"
	"monografia/transport/entity"

	"github.com/go-chi/chi/v5"
)

func NewRouter(srv service.Service, entity *entity.Entity) *chi.Mux {
	r := chi.NewRouter()

	orders := &Orders{service: &srv, entity: entity}
	products := &Products{service: &srv, entity: entity}

	// Orders
	r.Get("/users/{userID:[1-9][0-9]*}/orders", orders.GetByUserID)
	r.Get("/orders/{orderID:[1-9][0-9]*}", orders.GetByID)
	r.Post("/orders", orders.Create)
	r.Post("/orders/{orderID:[1-9][0-9]*}/item", orders.AddItem)
	r.Delete("/orders/item/{itemID:[1-9][0-9]*}", orders.RemoveItem)
	r.Put("/orders/{orderID:[1-9][0-9]*}/pay", orders.Pay)

	// Products
	r.Get("/products/{productID:[1-9][0-9]*}", products.GetByID)
	r.Post("/products", products.Create)

	return r
}
