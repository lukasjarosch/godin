package domain

import "errors"

// Domain errors
// These can then be remapped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
    ErrNotImplemented = errors.New("endpoint not implemented")
)
