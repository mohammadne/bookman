package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
)

type mysql struct {
	// injected dependencies
	config *Config

	// internal dependencies
	connection *sql.DB
}

const (
	driver = "mysql"

	errOpenDatabse = "error in opening mysql database"
	errPingDatabse = "error to ping mysql databse"
)

var (
	failurePrepareStatement = failures.Database{}.NewInternalServer("error when tying to prepare statement")
)

func NewMysqlDatabase(cfg *Config) Database {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema,
	)

	client, err := sql.Open(driver, dataSourceName)
	if err != nil {
		log.Fatal(errOpenDatabse, logger.Error(err))
		return nil
	}

	if err = client.Ping(); err != nil {
		log.Fatal(errPingDatabse, logger.Error(err))
		return nil
	}

	return &mysql{config: cfg, connection: client}
}

var (
	failureCreate    = failures.Database{}.NewInternalServer("error when tying to create entry")
	failureDuplicate = failures.Database{}.NewBadRequest("entry exists")
)

func (db *mysql) Create(query string, args []interface{}) (int64, failures.Failure) {
	stmt, err := db.connection.Prepare(query)
	if err != nil {
		return 0, failurePrepareStatement
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(args)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, failureDuplicate
		}

		return 0, failureCreate
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return 0, failureCreate
	}

	return id, nil
}

var (
	failureRead         = failures.Database{}.NewInternalServer("error when tying to read entry")
	failureReadNotFound = failures.Database{}.NewInternalServer("there is no entry with provided arguments")
)

func (db *mysql) Read(query string, args []interface{}, dest ...interface{}) failures.Failure {
	stmt, err := db.connection.Prepare(query)
	if err != nil {
		return failurePrepareStatement
	}
	defer stmt.Close()

	result := stmt.QueryRow(args)
	err = result.Scan(dest)
	if err != nil {
		if err == sql.ErrNoRows {
			return failureReadNotFound
		}

		return failureRead
	}

	return nil
}

var (
	failureUpdate = failures.Database{}.NewInternalServer("error when tying to update entry")
)

func (db *mysql) Update(query string, args []interface{}) failures.Failure {
	stmt, err := db.connection.Prepare(query)
	if err != nil {
		return failurePrepareStatement
	}
	defer stmt.Close()

	if _, err = stmt.Exec(args); err != nil {
		return failureUpdate
	}

	return nil
}

var (
	failureDelete = failures.Database{}.NewInternalServer("error when tying to delete entry")
)

func (db *mysql) Delete(query string, args []interface{}) failures.Failure {
	stmt, err := db.connection.Prepare(query)
	if err != nil {
		return failurePrepareStatement
	}
	defer stmt.Close()

	if _, err = stmt.Exec(args); err != nil {
		return failureDelete
	}

	return nil
}
