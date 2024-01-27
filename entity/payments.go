package entity

import "time"

type Payment struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	Amount          int       `json:"amount"`
	Token           string    `json:"token"`
	RedirectURL     string    `json:"redirect_url"`
	TransactionTime time.Time `json:"transaction_time"`
}
