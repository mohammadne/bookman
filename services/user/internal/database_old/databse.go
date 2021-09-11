package database_old

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
)

type Database interface {
	CreateUser(user *models.User) failures.Failure
	FindUserById(id int64) (*models.User, failures.Failure)
	FindUserByEmailAndPassword(email string, password string) (*models.User, failures.Failure)
	FindUserByEmail(email string) (*models.User, failures.Failure)
	UpdateUser(user *models.User) failures.Failure
	DeleteUser(user *models.User) failures.Failure
}

type mysql struct {
	// injected dependencies
	config *Config
	logger logger.Logger

	// internal dependencies
	connection *sql.DB
}

const (
	driver = "mysql"

	errOpenDatabse = "error in opening mysql database"
	errPingDatabse = "error to ping mysql databse"
)

func NewMysqlDatabase(cfg *Config, log logger.Logger) Database {
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

	return &mysql{config: cfg, logger: log, connection: client}
}
