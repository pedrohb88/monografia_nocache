package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"monografia/database"
	srv "monografia/service"
	"monografia/store/invoices"
	"monografia/store/payments"
	"monografia/transport"
)

func main() {

	// Database
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	// Stores
	paymentsStore := payments.New(&db)
	invoicesStore := invoices.New(&db)

	// Services
	service := srv.New(paymentsStore, invoicesStore)

	// Transport
	router := transport.NewRouter(service)

	log.Default().Println("Running server on port :3334")
	http.ListenAndServe(":3334", router)
}
