package models

import "errors"

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
	ErrNotFound     = errors.New("not found")
)
