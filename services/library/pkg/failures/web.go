package failures

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type webFailure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *webFailure) Message() string {
	return f.FailureMessage
}

func (f *webFailure) Status() int {
	return f.FailureStatus
}

func (f *webFailure) Causes() []string {
	return f.FailureCauses
}

func (f *webFailure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Web struct{}

func (h Web) New(message string, status int, causes []string) Failure {
	return &webFailure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (h Web) NewFromBytes(bytes []byte) (Failure, error) {
	failure := new(webFailure)
	if err := json.Unmarshal(bytes, failure); err != nil {
		return nil, errors.New("invalid json")
	}

	return failure, nil
}

func (h Web) NewBadRequest(message string) Failure {
	return &webFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusBadRequest,
	}
}

func (h Web) NewNotFound(message string) Failure {
	return &webFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusNotFound,
	}
}

func (h Web) NewUnauthorized(message string) Failure {
	return &webFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnauthorized,
	}
}

func (h Web) NewUnprocessableEntity(message string) Failure {
	return &webFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnprocessableEntity,
	}
}

func (h Web) NewNotImplemented() Failure {
	return &webFailure{
		FailureMessage: "not implemented",
		FailureStatus:  http.StatusNotImplemented,
	}
}

func (h Web) NewInternalServer(message string, errors ...error) Failure {
	causes := make([]string, 0, len(errors))
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &webFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
