package model

import (
	"time"
)

type UserAccount struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	CardNumber string    `json:"cardNumber"`
	Balance    float64   `json:"currentBalance"`
	CreatedAt  time.Time `json:"createdAt"`
}

type NewUserAccountBalance struct {
	AccountID           uint    `json:"accountID"`
	Name                string  `json:"name"`
	CardNumber          string  `json:"cardNumber"`
	ReplenishmentAmount float64 `json:"replenishmentAmount"`
}
