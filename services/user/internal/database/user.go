package database

import (
	"github.com/mohammadne/bookman/core/errors"
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/models"
)

type UserDatabase interface {
	CreateUser(user *models.User) *errors.DbError
	ReadUserById(id int64) (*models.User, *errors.DbError)
	ReadUserByEmailAndPassword(email string, password string) (*models.User, *errors.DbError)
	UpdateUser(old *models.User, new *models.User) *errors.DbError
	DeleteUser(user *models.User) *errors.DbError
}

const (
	queryCreateUser                 = "INSERT INTO users(first_name, last_name, email, date_created, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryReadUserById               = "SELECT id, first_name, last_name, email, date_created, FROM users WHERE id=?;"
	queryReadUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, FROM users WHERE email=? AND password=?"
	queryUpdateUser                 = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser                 = "DELETE FROM users WHERE id=?;"
)

var (
	errCreateUser = errors.NewDatabaseError("error when tying to create user", "database error")
)

func (db *mysql) CreateUser(user *models.User) *errors.DbError {
	stmt, err := db.connection.Prepare(queryCreateUser)
	if err != nil {
		(*db.logger).Error("error when trying to prepare create user statement", logger.Error(err))
		return errCreateUser
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
	if err != nil {
		(*db.logger).Error("error when trying to create user", logger.Error(err))
		return errCreateUser
	}

	user.Id, err = insertResult.LastInsertId()
	if err != nil {
		(*db.logger).Error("error when trying to get last insert id after creating a new user", logger.Error(err))
		return errCreateUser
	}

	return errors.NewNotImplementedDatabaseError()
}

func (db *mysql) ReadUserById(id int64) (*models.User, *errors.DbError) {
	return nil, errors.NewNotImplementedDatabaseError()
}

func (db *mysql) ReadUserByEmailAndPassword(e string, p string) (*models.User, *errors.DbError) {
	return nil, errors.NewNotImplementedDatabaseError()
}

func (db *mysql) UpdateUser(old *models.User, new *models.User) *errors.DbError {
	return errors.NewNotImplementedDatabaseError()
}

func (db *mysql) DeleteUser(user *models.User) *errors.DbError {
	return errors.NewNotImplementedDatabaseError()
}
