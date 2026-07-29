package repository

import "errors"

var ErrInvalidRule = errors.New(
	"invalid lifecycle rule",
)
