package utils

import "net/http"

type RestError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewBadRequest(message string) *RestError {
	return &RestError{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Message: message,
	}
}
