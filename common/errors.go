package common

import "fmt"

// APIError represents an API error response
type APIError struct {
	StatusCode int
	Message    string
	URL        string
	Method     string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %s %s returned %d: %s",
		e.Method, e.URL, e.StatusCode, e.Message)
}

// IsNotFound checks if the error is a 404 Not Found error
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsConflict checks if the error is a 409 Conflict error
func IsConflict(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 409
	}
	return false
}

// IsUnauthorized checks if the error is a 401 Unauthorized error
func IsUnauthorized(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsForbidden checks if the error is a 403 Forbidden error
func IsForbidden(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 403
	}
	return false
}

// IsBadRequest checks if the error is a 400 Bad Request error
func IsBadRequest(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 400
	}
	return false
}

// IsServerError checks if the error is a 5xx Server Error
func IsServerError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode >= 500 && apiErr.StatusCode < 600
	}
	return false
}

// AuthError represents an authentication error
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("authentication error: %s", e.Message)
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}
