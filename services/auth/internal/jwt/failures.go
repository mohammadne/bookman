package jwt

import (
	"fmt"
	"net/http"

	"github.com/mohammadne/bookman/auth/pkg/failures"
)

type failure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *failure) Message() string {
	return f.FailureMessage
}

func (f *failure) Status() int {
	return f.FailureStatus
}

func (f *failure) Causes() []string {
	return f.FailureCauses
}

func (f *failure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Failure struct{}

func (Failure) New(message string, status int, causes []string) failures.Failure {
	return &failure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (Failure) NewBadRequest(message string) failures.Failure {
	return &failure{
		FailureMessage: message,
		FailureStatus:  http.StatusBadRequest,
	}
}

func (Failure) NewNotFound(message string) failures.Failure {
	return &failure{
		FailureMessage: message,
		FailureStatus:  http.StatusNotFound,
	}
}

func (Failure) NewUnauthorized(message string) failures.Failure {
	return &failure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnauthorized,
	}
}

func (Failure) NewUnprocessableEntity(message string) failures.Failure {
	return &failure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnprocessableEntity,
	}
}

func (Failure) NewNotImplemented() failures.Failure {
	return &failure{
		FailureMessage: "not implemented",
		FailureStatus:  http.StatusNotImplemented,
	}
}

func (Failure) NewInternalServer(message string, errors ...error) failures.Failure {
	causes := make([]string, 0, len(errors))
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &failure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
