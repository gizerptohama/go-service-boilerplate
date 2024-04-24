package utils

import (
	"boilerplate/internal/constants"
	"boilerplate/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateToken(uint, bool) (string, error)
}

type JWTTokenGenerator struct {
	secret any
}

func NewJWTTokenGenerator(secret any) (gen TokenGenerator) {
	return &JWTTokenGenerator{secret}
}

func (gen *JWTTokenGenerator) GenerateToken(userId uint, isAdmin bool) (string, error) {
	claims := entity.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(constants.JWT_Expire_Duration)),
			Issuer:    constants.JWT_Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		UserId:  userId,
		IsAdmin: isAdmin,
	}

	unsignedToken := jwt.NewWithClaims(
		constants.JWT_Sign_Method,
		claims,
	)

	token, err := unsignedToken.SignedString(gen.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func HashPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(hash)
	return
}
