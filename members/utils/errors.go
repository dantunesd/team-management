package utils

import "net/http"

type CustomError struct {
	message string
	code    int
}

func (d *CustomError) Error() string {
	return d.message
}

func (d *CustomError) Code() int {
	return d.code
}

func NewError(message string, code int) *CustomError {
	return &CustomError{
		message: message,
		code:    code,
	}
}

func NewBadRequest(message string) *CustomError {
	return NewError(message, http.StatusBadRequest)
}

func NewNotFound(message string) *CustomError {
	return NewError(message, http.StatusNotFound)
}
