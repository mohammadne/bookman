package failures

import (
	"fmt"
	"net/http"
)

type databaseFailure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *databaseFailure) Message() string {
	return f.FailureMessage
}

func (f *databaseFailure) Status() int {
	return f.FailureStatus
}

func (f *databaseFailure) Causes() []string {
	return f.FailureCauses
}

func (f *databaseFailure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Database struct{}

func (d Database) New(message string, status int, causes []string) Failure {
	return &databaseFailure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (d Database) NewNotFound(message string) Failure {
	return &databaseFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusNotFound,
	}
}

func (d Database) NewNotImplemented() Failure {
	return &databaseFailure{
		FailureMessage: "NOT IMPLEMENTED",
		FailureStatus:  http.StatusNotImplemented,
	}

}

func (d Database) NewInternalServer(message string, errors ...error) Failure {
	causes := make([]string, len(errors), 0)
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &databaseFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
