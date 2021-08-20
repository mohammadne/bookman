package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestError(message string, status int, err string, causes []interface{}) *RestError {
	return &RestError{
		Message: message,
		Status:  status,
		Error:   err,
		Causes:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (*RestError, error) {
	var apiErr *RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(message string, errors ...error) *RestError {
	causes := make([]interface{}, len(errors), 0)
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Causes:  causes,
	}
}
