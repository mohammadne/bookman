package database

import (
	"database/sql"
	"fmt"

	"github.com/mohammadne/bookman/core/logger"
)

type Database interface {
	UserDatabase
}

type mysql struct {
	// injected dependencies
	config *Config
	logger *logger.Logger

	// internal dependencies
	connection *sql.DB
}

const (
	driver = "mysql"

	errOpenDatabse = "error in opening mysql database"
	errPingDatabse = "error to ping mysql databse"
)

func NewMysqlDatabase(cfg *Config, log *logger.Logger) Database {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema,
	)

	client, err := sql.Open(driver, dataSourceName)
	if err != nil {
		(*log).Fatal(errOpenDatabse, logger.Error(err))
		return nil
	}

	if err = client.Ping(); err != nil {
		(*log).Fatal(errPingDatabse, logger.Error(err))
		return nil
	}

	return &mysql{config: cfg, logger: log, connection: client}
}
