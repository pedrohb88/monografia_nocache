package model

type Order struct {
	ID            int     `json:"id"                db:"id"`
	UserID        int     `json:"user_id"           db:"user_id"`
	ItemsQuantity int     `json:"items_quantity"    db:"items_quantity"`
	Price         float64 `json:"price"             db:"price"`
	PaymentID     *int    `json:"payment_id"        db:"payment_id"`
}
