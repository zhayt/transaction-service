package model

import "time"

type Transaction struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"userID"`
	Card      string  `json:"card"`
	Amount    float64 `json:"amount"`
	Items     []*TransactionItem
	CreatedAt time.Time `json:"createdAt"`
}

type TransactionItem struct {
	ID          uint `json:"ID"`
	Transaction *Transaction
	ProductName string  `json:"ProductName"`
	Price       float64 `json:"Price"`
	Quantity    int     `json:"quantity"`
}
