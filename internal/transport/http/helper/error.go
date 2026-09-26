package helper

import (
	"errors"
)

var (
	ErrValidation 			= errors.New("validation error")
	ErrBadParam 			= errors.New("bad parameter")
	ErrUnauthorized 		= errors.New("unauthorized")
	ErrForbidden 			= errors.New("forbidden")
)

type Error struct {
    Message string
    Cause   error
}

func (e *Error) Error() string {
    return e.Message
}

func (e *Error) Unwrap() error {
    return e.Cause
}

func NewError(cause error, message string) error {
    return &Error{
        Cause:   cause,
        Message: message,
    }
}


type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Fields []FieldError
}

func (e ValidationError) Error() string {
	return ErrValidation.Error()
}

func (e ValidationError) Unwrap() error {
	return ErrValidation
}

func NewValidationError(fields ...FieldError) ValidationError {
	return ValidationError{Fields: fields}
}

func ValidationFields(err error) ([]FieldError, bool) {
	if validationErr, ok := errors.AsType[ValidationError](err); ok {
		return validationErr.Fields, true
	}

	return nil, false
}
