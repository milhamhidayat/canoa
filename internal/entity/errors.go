package entity

import (
	"errors"
	"fmt"
)

var (
	// ErrNotFound is an error message when a resource is not found
	ErrNotFound = errors.New("resource is not found")
)

// ConstraintError is a named type for custom error on contstraint things.
type ConstraintError string

func (e ConstraintError) Error() string {
	return string(e)
}

// ConstraintErrorf creates new interface value of type error
func ConstraintErrorf(text string, value ...interface{}) ConstraintError {
	return ConstraintError(fmt.Sprintf(text, value...))
}
