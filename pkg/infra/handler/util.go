package handler

import (
	"encoding/json"
	"log"
)

// errorResponse represents http response when error occurred
type errorResponse struct {
	Message string `json:"message"`
}

// toErrorResponse convert from error to errorResponse
func toErrorResponse(err error) *errorResponse {
	return &errorResponse{Message: err.Error()}
}

func logWithJson(prefix string, data interface{}) {
	contents, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal data for log: %v", err)
		return
	}
	log.Println(prefix + ": " + string(contents))
}
