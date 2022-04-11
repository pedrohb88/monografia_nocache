package errors

import (
	"database/sql"
	"errors"
)

var (
	ErrInvoiceNotFound = errors.New("invoice not found")
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, sql.ErrNoRows)
}
