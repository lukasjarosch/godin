package template

import (
	"github.com/sirupsen/logrus"
)

type Template interface {
	Render() error
}

type template struct {
}

func NewTemplate() *template {
	return &template{}
}

func (t *template) Render() error {
	logrus.Info("called render")
	return nil
}
