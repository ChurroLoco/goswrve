package goswrve

import (
	"fmt"
)

// SwrveError represents API errors
type APIError struct {
	Code    int
	Message string
}

// Error adheres to hte Error interface by returning a string.
func (e *APIError) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

// CreateAPIError creates a pointer to a populated new APIError object.
func CreateAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
