package domain

import "errors"

var (
	ErrConflict 			= errors.New("conflict")
	ErrNotFound 			= errors.New("not found")
	ErrDuplicate 			= errors.New("duplicate entry")
	ErrInternal 			= errors.New("internal app error")
)

type AppError struct {
    Message string
    Cause   error
}

func (e *AppError) Error() string {
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Cause
}

func NewAppError(cause error, msg string) error {
	return &AppError{
		Cause: cause,
		Message: msg,
	}
}
