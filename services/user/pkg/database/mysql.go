package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type mysql struct {
	connection *sql.DB
}

const (
	driver = "mysql"

	errOpenDatabse = "error in opening mysql database"
	errPingDatabse = "error to ping mysql databse"
)

const (
	errPrepareStatement = "error when tying to prepare statement"
)

func NewMysql(cfg *Config) Database {
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

	return &mysql{connection: client}
}

var (
	failureCreate    = failures.Database{}.NewInternalServer("error when tying to create entry")
	failureDuplicate = failures.Database{}.NewBadRequest("entry exists")
)

func (db *mysql) Create(query string, args []interface{}) (int64, failures.Failure) {
	stmt, err := db.connection.Prepare(query)
	if err != nil {
		return 0, failures.Database{}.NewInternalServer(errPrepareStatement, err)
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
		return failures.Database{}.NewInternalServer(errPrepareStatement, err)
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
		return failures.Database{}.NewInternalServer(errPrepareStatement, err)
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
		return failures.Database{}.NewInternalServer(errPrepareStatement, err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(args); err != nil {
		return failureDelete
	}

	return nil
}
