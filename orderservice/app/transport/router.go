package transport

import (
	"monografia/service"

	"github.com/go-chi/chi/v5"
)

func NewRouter(srv service.Service) *chi.Mux {
	r := chi.NewRouter()

	orders := &Orders{service: &srv}
	products := &Products{service: &srv}

	// Orders
	r.Get("/users/{userID:[1-9][0-9]*}/orders", orders.GetByUserID)
	r.Get("/orders/{orderID:[1-9][0-9]*}", orders.GetByID)
	r.Post("/orders", orders.Create)
	r.Post("/orders/{orderID:[1-9][0-9]*}/items", orders.AddItem)
	r.Delete("/orders/{orderID:[1-9][0-9]*}/items/{itemID:[1-9][0-9]*}", orders.RemoveItem)
	r.Put("/orders/{orderID:[1-9][0-9]*}/pay", orders.Pay)

	// Products
	r.Get("/products", products.GetAll)
	r.Get("/products/{productID:[1-9][0-9]*}", products.GetByID)
	r.Post("/products", products.Create)
	r.Delete("/products/{productID:[1-9][0-9]*}", products.Delete)

	return r
}
