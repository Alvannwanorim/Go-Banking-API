package main

import "math/rand"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber int64  `json:"account_number"`
	Balance       int64  `json:"balance"`
}

func NewAccount(first_name, last_name string) *Account {
	return &Account{
		ID:            rand.Intn(1000),
		FirstName:     first_name,
		LastName:      last_name,
		AccountNumber: int64(rand.Intn(100000)),
	}
}
