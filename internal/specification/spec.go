package specification

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Specification holds everything required to construct a microservice and it's dependencies
type Specification struct {
	Project              Project
	Service              Service
	Models               Models
	ResolvedDependencies []ResolvedDependency
}

var (
	ErrModelUnspecified  = errors.New("UNSPECIFIED MODEL")
	ErrMethodUnspecified = errors.New("UNSPECIFIED METHOD")
)

func LoadPath(path string) (*Specification, error) {
	spec := &Specification{}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(raw, spec)
	if err != nil {
		return nil, err
	}

	return spec, nil
}

func (s *Specification) ResolveDependencies() error {
	for _, dep := range s.Service.Dependencies {
		resolved, err := dep.Resolve(s)
		if err != nil {
			return err
		}

		s.ResolvedDependencies = append(s.ResolvedDependencies, resolved)
	}
	return nil
}

func (s *Specification) Validate() error {
	for _, e := range s.Service.Errors {
		if err := e.Validate(); err != nil {
			return err
		}
	}
	return nil
}
