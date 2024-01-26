package entity

import "time"

type History struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	Amount          int       `json:"amount"`
	RedirectURL     string    `json:"redirect_url"`
	TransactionTime time.Time `json:"transaction_time"`
}
