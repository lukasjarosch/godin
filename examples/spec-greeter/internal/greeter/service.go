package greeter

import (
	"context"
	"errors"

	"github.com/lukasjarosch/godin/example/spec-greeter/internal/config"
	"github.com/sirupsen/logrus"
) // greeterAPI is the actual business-logic which you want to provide
type greeterAPI struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("the given name is empty")
)

// NewExampleAPI returns our business-implementation of the ExampleAPI
func NewGreeterAPI(config *config.Config, logger *logrus.Logger) *greeterAPI {

	service := &greeterAPI{
		logger: logger,
		config: config,
	}

	return service
}

// Greeting implements the business-logic for this RPC
func (svc *greeterAPI) Hello(ctx context.Context, name string) (greeting string, err error) {
}
