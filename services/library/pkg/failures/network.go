package failures

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type networkFailure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *networkFailure) Message() string {
	return f.FailureMessage
}

func (f *networkFailure) Status() int {
	return f.FailureStatus
}

func (f *networkFailure) Causes() []string {
	return f.FailureCauses
}

func (f *networkFailure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Network struct{}

func (Network) New(message string, status int, causes []string) Failure {
	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (Network) NewFromBytes(bytes []byte) (Failure, error) {
	failure := new(networkFailure)
	if err := json.Unmarshal(bytes, failure); err != nil {
		return nil, errors.New("invalid json")
	}

	return failure, nil
}

func (Network) NewBadRequest(message string) Failure {
	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusBadRequest,
	}
}

func (Network) NewNotFound(message string) Failure {
	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusNotFound,
	}
}

func (Network) NewUnauthorized(message string) Failure {
	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnauthorized,
	}
}

func (Network) NewUnprocessableEntity(message string) Failure {
	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnprocessableEntity,
	}
}

func (Network) NewNotImplemented() Failure {
	return &networkFailure{
		FailureMessage: "not implemented",
		FailureStatus:  http.StatusNotImplemented,
	}
}

func (Network) NewInternalServer(message string, errors ...error) Failure {
	causes := make([]string, 0, len(errors))
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &networkFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
