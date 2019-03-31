package greeter

import (
	"context"
	"errors"

	"github.com/lukasjarosch/godin/example/spec-greeter/internal/config"
	"github.com/sirupsen/logrus"
)

// Greeter is able to greet people
type Greeter struct {
	logger *logrus.Logger
	config config.Config
}

var (
	ErrEmptyName = errors.New("no name given, duh")
)

// NewGreeterAPI returns the business implementation of godin.greeter.v1beta1.GreeterAPI
func NewGreeterAPI(logger *logrus.Logger, config config.Config) *Greeter {

	service := &Greeter{
		logger: logger,
		config: config,
	}

	return service
}

// Hello implements the Hello() gRPC method
func (svc *Greeter) Hello(ctx context.Context, name string) (greeting Greeting, err error) {
	return Greeting{}, nil
}
