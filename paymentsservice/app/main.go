package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "myuser:mypass@tcp(db:3306)/test_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error opening connection: ", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error pinging connection: ", err)
		return
	}

	rows, err := db.Query("SELECT id, amount FROM payments")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var amount float64

		err := rows.Scan(&id, &amount)
		if err != nil {
			panic(err)
		}

		fmt.Println("result: ", id, amount)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
