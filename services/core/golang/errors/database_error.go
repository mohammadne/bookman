package errors

type DatabaseError struct{}

func NewNotImplementedDatabaseError() *DatabaseError {
	return &DatabaseError{}
}
