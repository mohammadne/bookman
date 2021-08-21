package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/core/errors"
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/models"
)

type Database interface {
	CreateUser(user *models.User) *errors.DbError
	ReadUserById(id int64) (*models.User, *errors.DbError)
	ReadUserByEmailAndPassword(email string, password string) (*models.User, *errors.DbError)
	UpdateUser(user *models.User) *errors.DbError
	DeleteUser(user *models.User) *errors.DbError
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
