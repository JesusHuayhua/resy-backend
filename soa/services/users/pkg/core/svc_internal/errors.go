package svc_internal

import "errors"

var (
	ErrUnknown         = errors.New("Unknown argument passed")
	ErrInvalidArgument = errors.New("Invalid argument passed")
)
