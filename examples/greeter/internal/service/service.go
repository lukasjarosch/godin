package service

import (
	"context"
	"errors"
)

// Yyy documentation is automatically
// added to the README.
type Yyy interface {
	// COMMENT
	Hello(ctx context.Context, name string) (greeting *Greeting, err error)
	// Comment irgendwas
	Hello2(ctx context.Context, name string) (greeting Greeting, err error)
	// Comment
	Hello3(ctx context.Context, name string) (greeting string, err error)
	// Comment
	Hello4(ctx context.Context, name []string) (greeting []Greeting, err error)
	// Comment
	Hello5(ctx context.Context, name []string) (greeting []*Greeting, err error)
	// Comment
	Hello6(ctx context.Context, name []*Greeting) (greeting []*Greeting, err error)
	// Comment
	Hello7(ctx context.Context, name *Greeting) (greeting []string, err error)
	// Comment
	Hello8(ctx context.Context, name *[]Greeting) (greeting []string, err error)
	// Comment
	Hello9(ctx context.Context, name *[]Greeting, foo string, bar string) (greeting []string, err error)
}

// Application errors
// These can then be remaped to transport-specific errors in the transport layer (gRPC, HTTP, AMQP ...)
var (
	ErrNotImplemented = errors.New("endpoint not implemented")
)

type Greeting struct{}
