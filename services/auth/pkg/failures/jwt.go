package failures

import (
	"fmt"
	"net/http"
)

type jwtFailure struct {
	FailureMessage string   `json:"message"`
	FailureStatus  int      `json:"status"`
	FailureCauses  []string `json:"causes"`
}

// ==============================================================> methods

func (f *jwtFailure) Message() string {
	return f.FailureMessage
}

func (f *jwtFailure) Status() int {
	return f.FailureStatus
}

func (f *jwtFailure) Causes() []string {
	return f.FailureCauses
}

func (f *jwtFailure) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - causes: %v",
		f.FailureMessage, f.FailureStatus, f.FailureCauses,
	)
}

// ==============================================================> constructors

type Jwt struct{}

func (Jwt) New(message string, status int, causes []string) Failure {
	return &jwtFailure{
		FailureMessage: message,
		FailureStatus:  status,
		FailureCauses:  causes,
	}
}

func (Jwt) NewUnprocessableToken(message string) Failure {
	return &jwtFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnprocessableEntity,
	}
}

func (Jwt) NewInvalid(message string) Failure {
	return &jwtFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusUnauthorized,
	}
}

func (Jwt) NewInternal(message string, errors ...error) Failure {
	causes := make([]string, 0, len(errors))
	for index := 0; index < len(errors); index++ {
		causes = append(causes, errors[index].Error())
	}

	return &jwtFailure{
		FailureMessage: message,
		FailureStatus:  http.StatusInternalServerError,
		FailureCauses:  causes,
	}
}
