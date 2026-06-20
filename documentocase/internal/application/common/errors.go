package application

import "errors"

var (
	ErrNaoAutorizado       = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
	ErrNotFound            = errors.New("NotFound")
	ErrConcurrencyConflict = errors.New("ErrConcurrencyConflict")
)
