package yyy

import (
	"context"
	"github.com/go-kit/kit/log"
	"yyy/internal/service"
)

type serviceImplementation struct {
	logger log.Logger
}

func NewServiceImplementation(logger log.Logger) *serviceImplementation {
	return &serviceImplementation{
		logger: logger,
	}
}

// COMMENT
func (s *serviceImplementation) Hello(ctx context.Context, name string) (greeting *service.Greeting, err error) {
}

// Comment irgendwas
func (s *serviceImplementation) Hello2(ctx context.Context, name string) (greeting service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello3(ctx context.Context, name string) (greeting string, err error) {
}

// Comment
func (s *serviceImplementation) Hello4(ctx context.Context, name []string) (greeting []service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello5(ctx context.Context, name []string) (greeting []*service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello6(ctx context.Context, name []*service.Greeting) (greeting []*service.Greeting, err error) {
}

// Comment
func (s *serviceImplementation) Hello7(ctx context.Context, name *service.Greeting) (greeting []string, err error) {
}

// Comment
func (s *serviceImplementation) Hello8(ctx context.Context, name *[]service.Greeting) (greeting []string, err error) {
}

// Comment
func (s *serviceImplementation) Hello9(ctx context.Context, name *[]service.Greeting, foo string, bar string) (greeting []string, err error) {
}
