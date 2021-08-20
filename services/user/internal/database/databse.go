package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/models"
)

type Database interface {
	Create(user *models.User) error
	ReadById(id int64) (*models.User, error)
	ReadByEmailAndPassword(email string, password string) (*models.User, error)
	Update(old *models.User, new *models.User) error
	Delete(user *models.User) error
}

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

	return &mysql{config: cfg, connection: client}
}

func (db *mysql) Create() error {
	return errors.New("NOT IMPLEMENTED")
}

func (db *mysql) ReadById(id int64) (*models.User, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

func (db *mysql) ReadByEmailAndPassword(email string, password string) (*models.User, error) {
	return nil, errors.New("NOT IMPLEMENTED")
}

func (db *mysql) Update(old *models.User, new *models.User) error {
	return errors.New("NOT IMPLEMENTED")
}

func (db *mysql) Delete(user *models.User) error {
	return errors.New("NOT IMPLEMENTED")
}
