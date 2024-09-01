package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

const INVALID_ID = "Invalid ID"
const INVALID_REQUEST_BODY = "Invalid Request Body"
const INVALID_EMAIL_OR_PASSWORD = "Invalid email or password"

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Details: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
