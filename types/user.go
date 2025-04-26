package types

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID    string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
