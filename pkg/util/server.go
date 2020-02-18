package util

type ErrorResponse struct {
	Message string
}

func ToErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
