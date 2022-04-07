package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"monografia/database"
	srv "monografia/service"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/products"
	"monografia/transport"
)

func main() {

	// Database
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	// Stores
	ordersStore := orders.New(&db)
	productsStore := products.New(&db)
	itemsStore := items.New(&db)

	// Services
	service := srv.New(ordersStore, productsStore, itemsStore)

	// Transport
	router := transport.NewRouter(service)

	log.Default().Println("Running server on port :3333")
	http.ListenAndServe(":3333", router)
}
