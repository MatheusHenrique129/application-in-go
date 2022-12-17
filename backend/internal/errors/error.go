package errors

import (
	"fmt"
	"net/http"

	"github.com/MatheusHenrique129/application-in-go/internal/consts"
)

type CauseList []interface{}

type CustomError interface {
	Message() string
	Code() string
	Status() int
	Cause() CauseList
	Error() string
}

type customError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (e customError) Code() string {
	return e.ErrorCode
}

func (e customError) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (e customError) Status() int {
	return e.ErrorStatus
}

func (e customError) Cause() CauseList {
	return e.ErrorCause
}

func (e customError) Message() string {
	return e.ErrorMessage
}

func NewBadRequestResponse(msg string) CustomError {
	return NewBadRequestCustomError(msg)
}

func NewValidationErrorResponse(errorList []interface{}) CustomError {
	return NewValidationCustomError(consts.ValidationErrorMessage, consts.ValidationErrorCode, errorList)
}

func NewValidationSingleErrorResponse(err interface{}) CustomError {
	return NewValidationErrorResponse([]interface{}{err})
}

func NewCustomError(message string, error string, status int, cause CauseList) CustomError {
	return customError{message, error, status, cause}
}

func NewNotFoundCustomError(message string) CustomError {
	return customError{message, "not_found", http.StatusNotFound, CauseList{}}
}

func NewTooManyRequestsCustomError(message string) CustomError {
	return customError{message, "too_many_requests", http.StatusTooManyRequests, CauseList{}}
}

func NewBadRequestCustomError(message string) CustomError {
	return customError{message, "bad_request", http.StatusBadRequest, CauseList{}}
}

func NewValidationCustomError(message string, error string, cause CauseList) CustomError {
	return customError{message, error, http.StatusBadRequest, cause}
}

func NewMethodNotAllowedCustomError() CustomError {
	return customError{"Method not allowed", "method_not_allowed", http.StatusMethodNotAllowed, CauseList{}}
}

func NewInternalServerCustomError(message string, err error) CustomError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return customError{message, "internal_server_error", http.StatusInternalServerError, cause}
}

func NewForbiddenCustomError(message string) CustomError {
	return customError{message, "forbidden", http.StatusForbidden, CauseList{}}
}

func NewUnauthorizedCustomError(message string) CustomError {
	return customError{message, "unauthorized_scopes", http.StatusUnauthorized, CauseList{}}
}

func NewConflictCustomError(id string) CustomError {
	return customError{"Can't update " + id + " due to a conflict error", "conflict_error", http.StatusConflict, CauseList{}}
}
