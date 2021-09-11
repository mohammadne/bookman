package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/library/internal/database/ent"
	impl "github.com/mohammadne/bookman/library/internal/database/impl"
)

type Database interface {
	impl.Author
	impl.Book
	impl.Migration
}

// NewClient creates an ent database connection based on entry config.
func NewClient(config *Config) (Database, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		config.User, config.Password, config.Host, config.DatabaseName,
	)

	client, err := ent.Open(config.Driver, dataSourceName)
	return impl.New(client), err
}
