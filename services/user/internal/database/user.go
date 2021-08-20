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
	UpdateUser(user *models.User) *errors.DbError
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
	errCreateUser                 = errors.NewDatabaseError("error when tying to create user")
	errReadUserById               = errors.NewDatabaseError("error when tying to read user")
	errReadUserByEmailAndPassword = errors.NewDatabaseError("error when tying to read user")
	errUpdateUser                 = errors.NewDatabaseError("error when tying to update user")
	errDeleteUser                 = errors.NewDatabaseError("error when tying to delete user")
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
	stmt, err := db.connection.Prepare(queryReadUserById)
	if err != nil {
		(*db.logger).Error("error when trying to prepare read user statement", logger.Error(err))
		return nil, errReadUserById
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(id)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		(*db.logger).Error("error when trying to read user by id", logger.Error(err))
		return nil, errReadUserById
	}

	return user, nil
}

func (db *mysql) ReadUserByEmailAndPassword(e string, p string) (*models.User, *errors.DbError) {
	stmt, err := db.connection.Prepare(queryReadUserByEmailAndPassword)
	if err != nil {
		(*db.logger).Error("error when trying to prepare read user statement", logger.Error(err))
		return nil, errReadUserByEmailAndPassword
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(e, p)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		(*db.logger).Error("error when trying to read user by email and password", logger.Error(err))
		return nil, errReadUserByEmailAndPassword
	}

	return user, nil
}

func (db *mysql) UpdateUser(user *models.User) *errors.DbError {
	stmt, err := db.connection.Prepare(queryUpdateUser)
	if err != nil {
		(*db.logger).Error("error when trying to prepare update user statement", logger.Error(err))
		return errUpdateUser
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		(*db.logger).Error("error when trying to update user", logger.Error(err))
		return errUpdateUser
	}

	return nil
}

func (db *mysql) DeleteUser(user *models.User) *errors.DbError {
	stmt, err := db.connection.Prepare(queryDeleteUser)
	if err != nil {
		(*db.logger).Error("error when trying to prepare delete user statement", logger.Error(err))
		return errDeleteUser
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		(*db.logger).Error("error when trying to delete user", logger.Error(err))
		return errDeleteUser
	}

	return nil
}
