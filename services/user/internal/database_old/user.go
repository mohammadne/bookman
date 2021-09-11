package database_old

import (
	"database/sql"
	"strings"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/utils"
	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
)

const (
	queryCreateUser                 = "INSERT INTO users(first_name, last_name, email, date_created, password) VALUES(?, ?, ?, ?, ?);"
	queryFindUserById               = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=? AND password=?"
	queryFindUserByEmail            = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=?"
	queryUpdateUser                 = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser                 = "DELETE FROM users WHERE id=?;"
)

var (
	failureCreateUser    = failures.Database{}.NewInternalServer("error when tying to create user")
	failureDuplicateUser = failures.Database{}.NewBadRequest("entry exists")

	failureFindUserById         = failures.Database{}.NewInternalServer("error when tying to read user")
	failureFindUserByIdNotFound = failures.Database{}.NewInternalServer("there is no user with requested ID")

	failureReadUserByEmailAndPassword = failures.Database{}.NewInternalServer("error when tying to read user")

	failureReadUserByEmail = failures.Database{}.NewInternalServer("error when tying to read user")

	failureUpdateUser = failures.Database{}.NewInternalServer("error when tying to update user")

	failureDeleteUser = failures.Database{}.NewInternalServer("error when tying to delete user")
)

func (db *mysql) CreateUser(user *models.User) failures.Failure {
	stmt, err := db.connection.Prepare(queryCreateUser)
	if err != nil {
		db.logger.Error("error when trying to prepare create user statement", logger.Error(err))
		return failureCreateUser
	}
	defer stmt.Close()

	user.DateCreated = utils.NowDatabseFormatString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return failureDuplicateUser
		}

		db.logger.Error("error when trying to create user", logger.Error(err))
		return failureCreateUser
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		db.logger.Error("error when trying to get last insert id after creating a new user", logger.Error(err))
		return failureCreateUser
	}

	user.Id = uint64(userId)
	return nil
}

func (db *mysql) FindUserById(id int64) (*models.User, failures.Failure) {
	stmt, err := db.connection.Prepare(queryFindUserById)
	if err != nil {
		db.logger.Error("error when trying to prepare read user statement", logger.Error(err))
		return nil, failureFindUserById
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(id)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, failureFindUserByIdNotFound
		}

		db.logger.Error("error when trying to read user by id", logger.Error(err))
		return nil, failureFindUserById
	}

	return user, nil
}

func (db *mysql) FindUserByEmailAndPassword(e string, p string) (*models.User, failures.Failure) {
	stmt, err := db.connection.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		db.logger.Error("error when trying to prepare read user statement", logger.Error(err))
		return nil, failureReadUserByEmailAndPassword
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(e, p)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		db.logger.Error("error when trying to read user by email and password", logger.Error(err))
		return nil, failureReadUserByEmailAndPassword
	}

	return user, nil
}

func (db *mysql) FindUserByEmail(email string) (*models.User, failures.Failure) {
	stmt, err := db.connection.Prepare(queryFindUserByEmail)
	if err != nil {
		db.logger.Error("error when trying to prepare read user statement", logger.Error(err))
		return nil, failureReadUserByEmail
	}
	defer stmt.Close()

	user := new(models.User)
	result := stmt.QueryRow(email)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		db.logger.Error("error when trying to read user by email", logger.Error(err))
		return nil, failureReadUserByEmail
	}

	return user, nil
}

func (db *mysql) UpdateUser(user *models.User) failures.Failure {
	stmt, err := db.connection.Prepare(queryUpdateUser)
	if err != nil {
		db.logger.Error("error when trying to prepare update user statement", logger.Error(err))
		return failureUpdateUser
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		db.logger.Error("error when trying to update user", logger.Error(err))
		return failureUpdateUser
	}

	return nil
}

func (db *mysql) DeleteUser(user *models.User) failures.Failure {
	stmt, err := db.connection.Prepare(queryDeleteUser)
	if err != nil {
		db.logger.Error("error when trying to prepare delete user statement", logger.Error(err))
		return failureDeleteUser
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		db.logger.Error("error when trying to delete user", logger.Error(err))
		return failureDeleteUser
	}

	return nil
}
