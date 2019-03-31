package greeter

import (
	"fmt"

	"errors"

	"github.com/lukasjarosch/godin/examples/greeter/internal/config"
	"github.com/sirupsen/logrus"
)

// ExampleAPI is the actual business-logic which you want to provide
type greeterAPI struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("the given name is empty")
)

// NewExampleAPI returns our business-implementation of the ExampleAPI
func NewExampleAPI(config *config.Config, logger *logrus.Logger) *ExampleAPI {

	service := &ExampleAPI{
		logger: logger,
		config: config,
	}

	return service
}

// Greeting implements the business-logic for this RPC
func (e *ExampleAPI) Hello(name string) (greeting string, err error) {
	if name == "" {
		return "", ErrEmptyName
	}

	return fmt.Sprintf("Hey there, " + name), nil
}
