package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/joho/godotenv"
)

func New() (gorp.DbMap, error) {

	if os.Getenv("ENV") != "invoiceion" {
		err := godotenv.Load()
		if err != nil {
			return gorp.DbMap{}, fmt.Errorf("error loading config; %w", err)
		}
	}

	// Database
	db, err := sql.Open("mysql", os.Getenv("DBCONN"))
	if err != nil {
		return gorp.DbMap{}, err
	}

	dbmap := gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap, nil
}
