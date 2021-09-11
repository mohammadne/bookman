package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	database_impl "github.com/mohammadne/bookman/library/internal/database/calls"
	"github.com/mohammadne/bookman/library/internal/database/ent"
)

const dataSourceSchema = "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s"

// NewClient creates an ent database connection based on entry config.
func NewClient(config *Config) (database_impl.Database, error) {
	dataSourceName := fmt.Sprintf(
		dataSourceSchema,
		config.Host, config.Port, config.User,
		config.DatabaseName, config.Password, config.SSLMode,
	)

	client, err := ent.Open(config.Driver, dataSourceName)
	return database_impl.New(client), err
}
