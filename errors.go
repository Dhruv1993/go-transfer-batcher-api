package app

import (
	"fmt"
	"net/http"
)

type ErrorCode string
type ErrorData map[string]interface{}

const (
	ErrCodeInternal   ErrorCode = "internal_error"
	ErrCodeBadRequest ErrorCode = "invalid_request"
	ErrCodeValidation ErrorCode = "VALIDATION_ERROR"
)

type Error struct {
	Status int       `json:"status"`
	Code   ErrorCode `json:"code"`
	Title  string    `json:"title"`
	Detail string    `json:"detail"`
	Data   ErrorData `json:"data,omitempty"`
	Err    error     `json:"-"`
	Stack  []byte    `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s. %s. Err=%s. Data=%+v", e.Code, e.Title, e.Detail, e.Err, e.Data)
}

func NewErr(status int, code ErrorCode, title, detail string, data ErrorData, err error) *Error {
	return &Error{
		Status: status,
		Code:   code,
		Title:  title,
		Detail: detail,
		Data:   data,
		Err:    err,
	}
}

// common errors
func BadRequest(err error) error {
	return NewErr(http.StatusBadRequest, ErrCodeBadRequest, "Invalid Request", "Invalid JSON in request body.", nil, err)
}

func ValidationFailed(data ErrorData) error {
	return NewErr(http.StatusUnprocessableEntity, ErrCodeValidation, "Data Validation Error", "One or more fields could not be validated.", data, nil)
}
