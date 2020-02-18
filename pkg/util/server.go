package util

// ErrorResponse represents http response when error occurred
type ErrorResponse struct {
	Message string
}

// ToErrorResponse convert from error to ErrorResponse
func ToErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
