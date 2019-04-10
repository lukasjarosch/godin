package greeter

import (
	"context"

	"github.com/lukasjarosch/godin/examples/spec-greeter/internal/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

// Greeter is able to greet people
type Greeter struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = status.Error(codes.InvalidArgument, "no name given, duh")
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

// Burp implements the Burp() gRPC method
func (svc *Greeter) Burp(ctx context.Context) (name string, err error) {
	return "", nil
}
func (svc *Greeter) Goodbye(ctx context.Context, name string) (goodbye string, err error) {
	return "", nil
}
