package database

import "github.com/mohammadne/bookman/user/pkg/failures"

type Database interface {
	Create(query string, args []interface{}) (uint64, failures.Failure)

	Read(query string, args []interface{}, dest ...interface{}) failures.Failure

	Update(query string, args []interface{}) failures.Failure

	Delete(query string, args []interface{}) failures.Failure
}
