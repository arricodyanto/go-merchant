package model

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	jwt.RegisteredClaims
	CustomerID string `json:"customer_id"`
	Contact    string `json:"contact"`
}
