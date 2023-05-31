package models

type Payment struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
	Method string  `json:"method"`
}
