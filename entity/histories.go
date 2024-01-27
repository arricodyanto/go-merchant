package entity

import "time"

type History struct {
	ID                string    `json:"id"`
	CustomerID        string    `json:"customer_id"`
	Amount            int       `json:"amount"`
	Currency          string    `json:"currency"`
	PaymentType       string    `json:"payment_type"`
	TransactionStatus string    `json:"transaction_status"`
	TransactionTime   time.Time `json:"transaction_time"`
	SettlementTime    time.Time `json:"settlement_time"`
	ExpiryTime        time.Time `json:"expiry_time"`
}
