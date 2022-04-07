package model

type Order struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	ItemsQuantity int     `json:"items_quantity"`
	Price         float64 `json:"price"`
	ItemID        *int
	ItemQuantity  *int
	ItemPrice     *float64
	ProductID     *int
	ProductName   *string
	ProductPrice  *float64
}
