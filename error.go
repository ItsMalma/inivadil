package inivadil

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidValidationType = errors.New("invalid validation type")
)

type ValidationError struct {
	Field   string
	Message string
	stop    bool
}

func NewValidationError(field, message string, args ...any) ValidationError {
	return ValidationError{Field: field, Message: fmt.Sprintf(message, args...)}
}

func NilValidationError() ValidationError {
	return ValidationError{Field: "", Message: ""}
}

func (err ValidationError) Nil() bool {
	return err.Field == "" && err.Message == ""
}

func (err ValidationError) Error() string {
	if err.Nil() {
		return ""
	}
	return fmt.Sprintf("%v: %v", err.Field, err.Message)
}

type ValidationErrors []ValidationError

func (err ValidationErrors) Error() string {
	builder := new(strings.Builder)
	i := 0
	for _, validationError := range err {
		if validationError.Nil() {
			continue
		}
		if i > 0 {
			builder.WriteByte('\n')
		}
		builder.WriteString(validationError.Error())
		i++
	}
	return builder.String()
}
