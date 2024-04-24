package errors

import (
	"errors"
)

var (
	ErrInvalidToken           = errors.New("invalid token")
	ErrNoAuthHeader           = errors.New("no authorization header")
	ErrInvalidAuthHeader      = errors.New("invalid authorization header")
	ErrPageNotFound           = errors.New("page not found")
	ErrInvalidEmailPass       = errors.New("invalid email or password")
	ErrUserNotFound           = errors.New("user not found")
	ErrWalletNotFound         = errors.New("wallet not found")
	ErrSourceOfFundNotFound   = errors.New("source of fund not found")
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrNotEnoughBalance       = errors.New("not enough balance")
	ErrForgotPasswordNotFound = errors.New("forgot password entry not found")
	ErrCodeInvalid            = errors.New("invalid code")
	ErrFileImageNotFound      = errors.New("file image not found")
	ErrFileType               = errors.New("file type not acceptable (accept:png,jpg,jpeg,webp)")
	ErrFileSize               = errors.New("file size too large, expected not more than 1 mb")
	ErrProductNotFound        = errors.New("product not found")
)
