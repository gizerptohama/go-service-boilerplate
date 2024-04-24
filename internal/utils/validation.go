package utils

import (
	apperrors "boilerplate/internal/errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func MultiError(ve validator.ValidationErrors) (errs []error) {
	for _, fe := range ve {
		errs = append(errs, &apperrors.ValidationError{
			Field:   strings.ToLower(fe.Field()),
			Message: TagToMessage(fe),
		})
	}
	return
}

func TagToMessage(fe validator.FieldError) (message string) {
	switch fe.Tag() {
	case "required":
		message = "This field is required"
		return
	case "email":
		message = "Invalid email"
		return
	case "min":
		if fe.Type().Kind() == reflect.String {
			message = fmt.Sprintf("%s has to be at least %s characters long", fe.Field(), fe.Param())
			return
		}
		message = fmt.Sprintf("%s has a minimum value of %s", fe.Field(), fe.Param())
		return
	case "max":
		if fe.Type().Kind() == reflect.String {
			message = fmt.Sprintf("%s has to be at maximum %s characters long", fe.Field(), fe.Param())
			return
		}
		message = fmt.Sprintf("%s has a maximum value of %s", fe.Field(), fe.Param())
		return
	case "oneof":
		message = fmt.Sprintf("%s can only be one of the following: %v", fe.Field(), strings.ReplaceAll(fe.Param(), " ", ", "))
		return
	default:
		message = fe.Error()
		return
	}
}
