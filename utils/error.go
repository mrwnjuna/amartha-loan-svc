package utils

type AppError struct {
	StatusCode int
	Message    string
	Details    string
}

// Factory methods for common error types
func NewBadRequestError(message, details string) *AppError {
	return &AppError{StatusCode: 400, Message: message, Details: details}
}

func NewNotFoundError(message, details string) *AppError {
	return &AppError{StatusCode: 404, Message: message, Details: details}
}

func NewInternalServerError(message, details string) *AppError {
	return &AppError{StatusCode: 500, Message: message, Details: details}
}
