package failures

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type httpFailure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *httpFailure) Message() string {
	return f.FailureMessage
}

func (f *httpFailure) Status() int {
	return f.FailureStatus
}

func (f *httpFailure) Causes() []string {
	return f.FailureCauses
}

func (f *httpFailure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Rest struct{}

func (h Rest) New(message string, status int, causes []string) Failure {
	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (h Rest) NewFromBytes(bytes []byte) (Failure, error) {
	failure := new(httpFailure)
	if err := json.Unmarshal(bytes, failure); err != nil {
		return nil, errors.New("invalid json")
	}

	return failure, nil
}

func (h Rest) NewBadRequest(message string) Failure {
	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusBadRequest,
	}
}

func (h Rest) NewNotFound(message string) Failure {
	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusNotFound,
	}
}

func (h Rest) NewUnauthorized(message string) Failure {
	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnauthorized,
	}
}

func (h Rest) NewUnprocessableEntity(message string) Failure {
	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnprocessableEntity,
	}
}

func (h Rest) NewNotImplemented() Failure {
	return &httpFailure{
		FailureMessage: "not implemented",
		FailureStatus:  http.StatusNotImplemented,
	}
}

func (h Rest) NewInternalServer(message string, errors ...error) Failure {
	causes := make([]string, len(errors), 0)
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &httpFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
