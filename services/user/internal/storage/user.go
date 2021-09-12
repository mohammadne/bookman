package storage

import (
	"context"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

const (
	queryCreateUser                 = "INSERT INTO users(first_name, last_name, email, date_created, password) VALUES(?, ?, ?, ?, ?);"
	queryFindUserById               = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=? AND password=?"
	queryFindUserByEmail            = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email=?"
	queryUpdateUser                 = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser                 = "DELETE FROM users WHERE id=?;"
)

func (storage *storage) CreateUser(ctx context.Context, user *models.User) failures.Failure {
	_, span := storage.tracer.Start(ctx, "storage.user.create_user")
	defer span.End()

	args := []interface{}{user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password}
	id, failure := storage.database.Create(queryCreateUser, args)
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return failure
	}

	user.Id = id
	return nil
}

func (storage *storage) FindUserById(ctx context.Context, id uint64) (*models.User, failures.Failure) {
	_, span := storage.tracer.Start(ctx, "storage.user.find_user_by_id")
	defer span.End()

	user := new(models.User)
	dest := []interface{}{&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated}
	failure := storage.database.Read(queryFindUserById, []interface{}{id}, dest)
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return nil, failure
	}

	return user, nil
}

func (storage *storage) FindUserByEmailAndPassword(ctx context.Context, email string, password string) (*models.User, failures.Failure) {
	_, span := storage.tracer.Start(ctx, "storage.user.find_user_by_email_password")
	defer span.End()

	user := new(models.User)
	dest := []interface{}{&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated}
	failure := storage.database.Read(queryFindUserByEmailAndPassword, []interface{}{email, password}, dest)
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return nil, failure
	}

	return user, nil
}

func (storage *storage) FindUserByEmail(ctx context.Context, email string) (*models.User, failures.Failure) {
	_, span := storage.tracer.Start(ctx, "storage.user.find_user_by_email")
	defer span.End()

	user := new(models.User)
	dest := []interface{}{&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated}
	failure := storage.database.Read(queryFindUserByEmail, []interface{}{email}, dest)
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return nil, failure
	}

	return user, nil
}

func (storage *storage) UpdateUser(ctx context.Context, user *models.User) failures.Failure {
	_, span := storage.tracer.Start(ctx, "storage.user.update_user")
	defer span.End()

	args := []interface{}{user.FirstName, user.LastName, user.Email, user.Id}
	failure := storage.database.Update(queryUpdateUser, args)
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return failure
	}

	return nil
}

func (storage *storage) DeleteUser(ctx context.Context, user *models.User) failures.Failure {
	_, span := storage.tracer.Start(ctx, "storage.user.delete_user")
	defer span.End()

	failure := storage.database.Delete(queryDeleteUser, []interface{}{user.Id})
	if failure != nil {
		storage.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return failure
	}

	return nil
}
