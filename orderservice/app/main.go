package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"monografia/database"
	srv "monografia/service"
	"monografia/store/items"
	"monografia/store/orders"
	"monografia/store/payments"
	"monografia/store/products"
	"monografia/transport"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("error loading config; %s", err.Error())
		}
	}

	// Database
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	// Stores
	ordersStore := orders.New(&db)
	productsStore := products.New(&db)
	itemsStore := items.New(&db)
	paymentsStore := payments.New()

	// Services
	service := srv.New(ordersStore, productsStore, itemsStore, paymentsStore)

	// Transport
	router := transport.NewRouter(service)

	log.Default().Println("Running server on port :3333")
	http.ListenAndServe(":3333", router)
}
