package module

import (
	"github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
    return &Endpoint{}
}

func (e *Endpoint) Execute() error {
	logrus.Info("new enpoint created")
	return nil
}

func (e *Endpoint) prompt() (string, error) {
	return "", nil
}

