package errors

type DbError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewDatabaseError(message string, err string) *DbError {
	return &DbError{
		Message: message,
		Error:   err,
	}
}

func NewNotImplementedDatabaseError() *DbError {
	return &DbError{}
}
