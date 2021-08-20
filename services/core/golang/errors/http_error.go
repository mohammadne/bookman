package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HttpError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewHttpError(message string, status int, err string, causes []interface{}) *HttpError {
	return &HttpError{
		Message: message,
		Status:  status,
		Error:   err,
		Causes:  causes,
	}
}

func NewFromBytesHttpError(bytes []byte) (*HttpError, error) {
	var apiErr *HttpError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestHttpError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundHttpError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedHttpError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewNotImplementedHttpError() *HttpError {
	return &HttpError{
		Message: "not implemented",
		Status:  http.StatusNotImplemented,
		Error:   "not_implemented",
	}
}

func NewInternalServerHttpError(message string, errors ...error) *HttpError {
	causes := make([]interface{}, len(errors), 0)
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &HttpError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Causes:  causes,
	}
}
