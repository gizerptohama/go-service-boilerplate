package errors

import "fmt"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationError(field, message string) (err error) {
	err = &ValidationError{field, message}
	return
}

func (err *ValidationError) Error() string {
	return fmt.Sprintf("Error on field '%s': %s", err.Field, err.Message)
}
