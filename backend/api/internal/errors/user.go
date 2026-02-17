package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrEmailAlreadyUsed  = errors.New("email already used")
	ErrPhoneAlreadyUsed  = errors.New("phone already used")
)
