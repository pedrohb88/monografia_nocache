package errors

import (
	"database/sql"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, sql.ErrNoRows)
}
