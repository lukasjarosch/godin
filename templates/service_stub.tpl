package service

import (
    "errors"
)

// {{ title .Service.Name }} documentation is automatically
// added to the README.
type {{ title .Service.Name }} interface {
    // Hello greets you. This comment is also automatically added to the README.
    // Also make sure that all parameters are named, Godin requires this information in order to work.
    Hello(ctx context.Context, name string) (greeting string, err error)
}

// Application errors
// These can then be remaped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
    ErrNotImplemented = errors.New("endpoint not implemented")
)
