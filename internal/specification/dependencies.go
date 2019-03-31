package specification

import (
	"errors"
	"fmt"
	"path"
)

var RegisteredDependencies = []string{"config", "logger"}

type Dependency struct {
	Type string
}

func (d Dependency) Resolve(spec *Specification) (ResolvedDependency, error) {
	switch d.Type {
	case "config":
		return 	&Configuration{
			modulePath: spec.Project.Module,
		}, nil
	case "logger":
		return &Logger{}, nil
	}

	return nil, errors.New(fmt.Sprintf("unresolved dependency: %s", d.Type))
}

type ResolvedDependency interface {
	Type() string
	Name() string
	Import() string
}


type Logger struct {
}

func (l Logger) Type() string {
	return "*logrus.Logger"
}

func (l Logger) Name() string {
	return "logger"
}

func (l Logger) Import() string {
	return "github.com/sirupsen/logrus"
}


type Configuration struct {
	modulePath string
}

func (c Configuration) Type() string {
	return "config.Config"
}

func (c Configuration) Name() string {
	return "config"
}

func (c Configuration) Import() string {
	return path.Join(c.modulePath, "internal", "config")
}