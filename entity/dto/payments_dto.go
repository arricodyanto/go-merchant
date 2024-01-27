package dto

type TransactionDetailDto struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type CreditCardDto struct {
	Secure bool `json:"secure"`
}

type SnapRequestDto struct {
	TransactionDetails TransactionDetailDto `json:"transaction_details"`
	CreditCard         CreditCardDto        `json:"credit_card"`
}

type SnapResponseDto struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
