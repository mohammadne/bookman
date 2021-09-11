package database

import "github.com/mohammadne/bookman/library/pkg/failures"

type Database interface {
	Create(query string, args []interface{}) (int64, failures.Failure)

	Read(query string, args []interface{}, dest ...interface{}) failures.Failure

	Update(query string, args []interface{}) failures.Failure

	Delete(query string, args []interface{}) failures.Failure
}
