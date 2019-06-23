package service

import (
	"context"
	"errors"
)

// User documentation is automatically
// added to the README.
type User interface {
	// Create will create a new user and return it.
	Create(ctx context.Context, username string, email string) (user *UserEntity, err error)
}

// Application errors
// These can then be remaped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
	ErrNotImplemented = errors.New("endpoint not implemented")
)

type UserEntity struct {
	ID    string
	Name  string
	Email string
}
