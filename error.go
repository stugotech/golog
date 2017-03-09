package golog

import (
	"fmt"
)

// loggedError marks an error as logged
type loggedError struct {
	err    error
	logged bool
	tag    string
}

// LoggedError indicates if an error has been logged
type LoggedError interface {
	error
	Logged() bool
}

// NewError creates a new logged error
func NewError(tag string, msg string, args ...interface{}) LoggedError {
	return &loggedError{
		tag: tag,
		err: fmt.Errorf(msg, args...),
	}
}

// WrapError creates a new logged error from an existing one
func WrapError(tag string, err error) (LoggedError, bool) {
	logged, ok := err.(LoggedError)
	if !ok {
		logged = &loggedError{
			tag: tag,
			err: err,
		}
	}
	return logged, ok
}

// Error gets the error message
func (e *loggedError) Error() string {
	e.logged = true
	return e.tag + ": " + e.err.Error()
}

// Logged returns true if the error has been logged
func (e *loggedError) Logged() bool {
	return e.logged
}
