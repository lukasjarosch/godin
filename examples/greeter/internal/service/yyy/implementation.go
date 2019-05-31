package yyy

import (
	"github.com/go-kit/kit/log"
)

type serviceImplementation struct {
	logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
	return &serviceImplementation{
		logger: logger,
	}
}

// Hello greets you. This comment is also automatically added to the README.
// Also make sure that all parameters are named, Godin requires this information in order to work.
func (s *serviceImplementation) Hello(ctx context.Context, name string) (greeting *service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello2(ctx context.Context, name string) (greeting service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello3(ctx context.Context, name string) (greeting string, err error) {
}

// Comment
func (s *serviceImplementation) Hello4(ctx context.Context, name []service.string) (greeting []service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello5(ctx context.Context, name []service.string) (greeting []*service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello6(ctx context.Context, name []*service.Greeting) (greeting []*service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello7(ctx context.Context, name *service.Greeting) (greeting []service.string, err error) {
}
