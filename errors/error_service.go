package errors

import (
	"fmt"
	"net/http"
)

type ErrorService struct{}

var instance *ErrorService

func NewErrorService() *ErrorService {
	if instance == nil {
		instance = &ErrorService{}
	}
	return instance
}

func (es *ErrorService) NotFound(resource string) *AppError {
	return NewAppError(http.StatusNotFound, fmt.Sprintf("resource not found: %s", resource), nil)
}

func (es *ErrorService) InternalServerError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "internal server error", err)
}

func (es *ErrorService) BadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message, nil)
}

func (es *ErrorService) Unauthorized(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, message, nil)
}

func (es *ErrorService) Forbidden(message string) *AppError {
	return NewAppError(http.StatusForbidden, message, nil)
}
