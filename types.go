package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	Amount    int `json:"amount"`
	AccountId int `json:"account_id"`
}
type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber int64     `json:"account_number"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewAccount(first_name, last_name string) *Account {
	return &Account{
		FirstName:     first_name,
		LastName:      last_name,
		AccountNumber: int64(rand.Intn(100000)),
		CreatedAt:     time.Now().UTC(),
	}
}
