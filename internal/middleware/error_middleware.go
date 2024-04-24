package middleware

import (
	"boilerplate/internal/dto"
	apperrors "boilerplate/internal/errors"
	"boilerplate/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var clientError = &apperrors.ClientError{}
var validationError = validator.ValidationErrors{}
var syntaxErr = &json.SyntaxError{}
var unmarshallErr = &json.UnmarshalTypeError{}

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()

		if err == nil {
			return
		}

		if errors.Is(err, context.DeadlineExceeded) {
			ctx.AbortWithStatusJSON(http.StatusGatewayTimeout, dto.NewErrorResponse(err))
			return
		}

		if errors.As(err.Err, &clientError) {
			err, _ := err.Err.(*apperrors.ClientError)
			ctx.AbortWithStatusJSON(err.GetCode(), dto.NewErrorResponse(err))
			return
		}

		if errors.As(err, &syntaxErr) || errors.As(err, &unmarshallErr) ||
			errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err))
			return
		}

		if errors.As(err.Err, &validationError) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.NewMultiErrorResponse(utils.MultiError(validationError)...))
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
			return
		}
	}
}
