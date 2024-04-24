package dto

import "boilerplate/internal/entity"

type RegisterResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (dto *RegisterResponse) FromUser(user entity.User) {
	dto.ID = user.ID
	dto.UserName = user.UserName
	dto.Email = user.Email
}
