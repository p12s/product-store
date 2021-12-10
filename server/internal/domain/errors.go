package domain

import "errors"

var (
	// ErrProductNotFound
	ErrProductNotFound = errors.New("product doesn't exists")
	// ErrProductAlreadyExists
	ErrProductAlreadyExists = errors.New("product with such name already exists")
)
