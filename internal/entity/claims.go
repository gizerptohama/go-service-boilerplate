package entity

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UserId  uint `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}
