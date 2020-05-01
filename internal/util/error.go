package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type UnauthorizedError struct {
	OriginalError error
}

func (ue *UnauthorizedError) Error() string {
	return "Unauthorized"
}

func (err *UnauthorizedError) WithContext(message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args...)
}

func NewUnauthorizedError(err ...error) *UnauthorizedError {
	if len(err) > 0 {
		return &UnauthorizedError{OriginalError: err[0]}
	}

	return &UnauthorizedError{}
}

type BadRequestError struct {
	OriginalError error
	error
}

func (bre *BadRequestError) Error() string {
	return "Bad Request"
}

func (err *BadRequestError) WithContext(message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args...)
}

func NewBadRequestError(err ...error) *BadRequestError {
	if len(err) > 0 {
		return &BadRequestError{OriginalError: err[0]}
	}

	return &BadRequestError{}
}

type DatabaseError struct {
	OriginalError error
}

func (de *DatabaseError) Error() string {
	if de.OriginalError != nil {
		return de.OriginalError.Error()
	}

	return "Database Error"
}

func (err *DatabaseError) WithContext(message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args...)
}

func NewDatabaseError(err ...error) *DatabaseError {
	if len(err) > 0 {
		return &DatabaseError{OriginalError: err[0]}
	}

	return &DatabaseError{}
}

type ValidationError struct {
	OriginalError error
	errorMessages []string
}

func (err *ValidationError) Error() string {
	if err.OriginalError != nil {
		validationErrors := err.OriginalError.(validator.ValidationErrors)
		for _, e := range validationErrors {
			errMessage := fmt.Sprintf("validation failed on field %s, condition: %s", e.Field(), e.Tag())

			if e.Param() != "" {
				errMessage = fmt.Sprintf("%v {%v} ", errMessage, e.Param())
			}

			if e.Value() != nil && e.Value() != "" {
				errMessage = fmt.Sprintf("%v, actual: %v", errMessage, e.Value())
			}

			err.errorMessages = append(err.errorMessages, errMessage)
		}
	}

	return strings.Join(err.errorMessages, ".\n")
}

func (err *ValidationError) WithContext(message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args...)
}

func NewValidationError(err ...error) *ValidationError {
	if len(err) > 0 {
		return &ValidationError{OriginalError: err[0]}
	}

	return &ValidationError{}
}

func WrapError(err error, message string, args ...interface{}) error {
	return errors.Wrapf(err, message, args...)
}

func NewError(message string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(message, args...))
}
