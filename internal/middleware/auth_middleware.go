package middleware

import (
	"boilerplate/internal/constants"
	"boilerplate/internal/entity"
	apperrors "boilerplate/internal/errors"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(isAdmin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Next()
		auth := ctx.GetHeader("Authorization")

		if auth == "" {

			err := apperrors.NewClientError(apperrors.ErrNoAuthHeader, http.StatusUnauthorized)
			err = ctx.Error(err)
			ctx.Abort()
			return
		}

		tokenSplit := strings.Split(auth, " ")
		if len(tokenSplit) != 2 {
			err := apperrors.NewClientError(apperrors.ErrInvalidAuthHeader, http.StatusUnauthorized)
			err = ctx.Error(err)
			ctx.Abort()
			return
		}

		claims := &entity.Claims{}
		token, err := jwt.ParseWithClaims(tokenSplit[1], claims, keyFunc)
		if err != nil {
			err = apperrors.NewClientError(apperrors.ErrInvalidToken, http.StatusUnauthorized)
			err = ctx.Error(err)
			ctx.Abort()
			return
		}

		if !token.Valid {
			err = apperrors.NewClientError(apperrors.ErrInvalidToken, http.StatusUnauthorized)
			err = ctx.Error(err)
			ctx.Abort()
			return
		}

		if claims.IsAdmin != isAdmin {
			err = apperrors.NewClientError(apperrors.ErrInvalidToken, http.StatusUnauthorized)
			err = ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx2 := context.WithValue(ctx.Request.Context(), "user_id", claims.UserId)
		ctx.Request = ctx.Request.WithContext(ctx2)
	}
}

func keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return constants.JWT_Secret, nil
}
