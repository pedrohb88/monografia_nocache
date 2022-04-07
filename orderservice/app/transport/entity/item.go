package entity

type Item struct {
	ID       int      `json:"id"`
	OrderID  int      `json:"order_id"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
	Product  *Product `json:"product"`
}
