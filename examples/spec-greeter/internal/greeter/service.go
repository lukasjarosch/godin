package greeter

import (
	"context"
	"errors"

	"github.com/lukasjarosch/godin/examples/spec-greeter/internal/config"
	"github.com/sirupsen/logrus"
)

// Greeter is able to greet people
type Greeter struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("no name given, duh")
)

// NewGreeter returns the business implementation of godin.greeter.v1beta1.GreeterAPI
func NewGreeter(config *config.Config, logger *logrus.Logger) *Greeter {

	service := &Greeter{
		config: config,
		logger: logger,
	}

	return service
}

// Hello implements the Hello() gRPC method
func (svc *Greeter) Hello(ctx context.Context, name string) (greeting Greeting, err error) {
	return Greeting{}, nil
}
