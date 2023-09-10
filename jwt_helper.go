package main

import (
	"fmt"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	AccountNumber int64 `json:"account_number"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signed method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func CreateToken(account *Account) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"expiresAt":      15000,
		"account_number": account.AccountNumber,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss, err := token.SignedString(secret)
	fmt.Printf("%v %v", ss, err)
	return token, nil
}
