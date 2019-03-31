package specification

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Specification holds everything required to construct a microservice and it's dependencies
type Specification struct {
	Project Project
	Service Service
	Models Models
}

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