package dto

import "boilerplate/internal/entity"

type RegisterRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (dto *RegisterRequest) ToUser() (user entity.User) {
	user = entity.User{
		UserName: dto.UserName,
		Email:    dto.Email,
		Password: dto.Password,
	}
	return
}
