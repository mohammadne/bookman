package errors

type DbError struct {
	Message string `json:"message"`
}

func NewDatabaseError(message string) *DbError {
	return &DbError{Message: message}
}

func NewNotImplementedDatabaseError() *DbError {
	return &DbError{Message: "NOT IMPLEMENTED"}
}
