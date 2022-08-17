package errs

import "errors"

var (
	ErrUrlNotFound error = errors.New("Short url not found")
	ErrUrlNotSaved error = errors.New("couldn't save due to db error")
)
