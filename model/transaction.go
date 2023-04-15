package model

import "time"

type Transaction struct {
	ID         uint    `json:"id"`
	AccountID  uint    `json:"accountID"`
	CardNumber string  `json:"cardNumber"`
	Amount     float64 `json:"amount"`
	Items      []*TransactionItem
	CreatedAt  time.Time `json:"createdAt"`
}

type TransactionItem struct {
	ID            uint    `json:"ID"`
	TransactionID uint    `json:"transactionID"`
	ProductName   string  `json:"ProductName"`
	Price         float64 `json:"Price"`
	Quantity      int     `json:"quantity"`
}
