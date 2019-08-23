package models

import (
	"errors"
)

var(
	ErrNotFound = errors.New("Requested item not found!")
)