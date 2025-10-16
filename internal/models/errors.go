package models

import (
	"errors"
)

var (
	ErroNoRecord = errors.New("models: no matching record found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)
