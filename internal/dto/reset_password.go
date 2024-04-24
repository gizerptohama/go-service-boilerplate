package dto

type ResetPassword struct {
	Code     string `json:"code" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
