package handler

import (
	"errors"
	"net/http"

	"boilerplate/internal/dto"
	apperrors "boilerplate/internal/errors"
	"boilerplate/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) (h *UserHandler) {
	return &UserHandler{u}
}

func (h *UserHandler) GetUserData(ctx *gin.Context) {
	id, _ := ctx.Get("user_id")

	result, err := h.userUsecase.GetUserDataById(ctx, id.(uint))

	if errors.Is(err, apperrors.ErrUserNotFound) {
		err = ctx.Error(apperrors.NewClientError(err))
		return
	}
	if err != nil {
		err = ctx.Error(err)
		return
	}

	var registerRespDTO dto.RegisterResponse
	registerRespDTO.FromUser(*result)

	ctx.JSON(http.StatusOK, dto.NewDataResponse(registerRespDTO))

}
