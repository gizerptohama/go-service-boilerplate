package handler

import (
	"errors"
	"net/http"

	"boilerplate/internal/dto"
	apperrors "boilerplate/internal/errors"
	"boilerplate/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUscase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) (h *AuthHandler) {
	return &AuthHandler{u}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	result, err := h.authUscase.Register(ctx, registerRequest)
	if errors.Is(err, apperrors.ErrEmailAlreadyRegistered) {
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

func (h *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	token, err := h.authUscase.Login(ctx, loginRequest)
	if errors.Is(err, apperrors.ErrInvalidEmailPass) {
		err = ctx.Error(apperrors.NewClientError(err))
		return
	}
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.NewTokenResponse(token))
}
