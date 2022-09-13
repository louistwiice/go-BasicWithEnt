package entity

import "errors"

var  ErrNotFound = errors.New("not found")

var ErrInvalidModel = errors.New("invalid model")

var ErrCanNotBeDeleted = errors.New("can not be deleted")

var ErrNullField = errors.New("some fields should not be null")

var ErrInvalidPassword = errors.New("password is empty or invalid")