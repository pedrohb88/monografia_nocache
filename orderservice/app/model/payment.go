package model

type Payment struct {
	ID      int      `json:"id"`
	Amount  float64  `json:"amount"`
	Invoice *Invoice `json:"invoice"`
}
