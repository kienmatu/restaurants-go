package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Key        string `json:"errorKey"`
	Log        string `json:"log"`
}

func NewFullErrorResponse(status int, root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: status,
		RootErr:    root,
		Message:    message,
		Key:        key,
		Log:        log,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return NewFullErrorResponse(http.StatusUnauthorized, root, msg, "", key)
}

func NewErrResponse(root error, msg, log, key string) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, root, msg, log, key)
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrResponse(root, msg, root.Error(), key)
	}
	return NewErrResponse(errors.New(msg), msg, msg, key)
}

/*
Okay, get error recursively
*/
func (a *AppError) Error() string {
	return a.RootError().Error()
}

/*
RootError
- Get root error recursively
*/
func (a *AppError) RootError() error {
	if err, ok := a.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return a.RootErr
}

// Errors

func ErrInvalidRequest(err error) *AppError {
	return NewErrResponse(err, "your request is invalid", err.Error(), "bad_request")
}

func ErrDB(err error) *AppError {
	return NewErrResponse(err, "error with DB", err.Error(), "db_error")
}

func ErrInternal(err error) *AppError {
	return NewErrResponse(err, "something went wrong with server", err.Error(), "ErrInternal")
}
func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("can not create %s", entity), // message
		fmt.Sprintf("ErrCannotCreate%s", entity), // key
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("can not update %s", entity),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("can not delete %s", entity),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("can not list %s", entity),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
