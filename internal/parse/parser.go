package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

type baseParser struct {
	File      *types.File
	Interface *types.Interface
	Path      string
}

// ParseFile will use go-astra to parse the service-File
func (p *baseParser) ParseFile() (err error) {
	p.File, err = astra.ParseFile(p.Path)
	if err != nil {
		return errors.Wrap(err, "ParseFile")
	}
	return nil
}

// FindInterface searches for an interface with the given name
func (p *baseParser) FindInterface(interfaceName string) (iface types.Interface, err error) {
	for _, iface := range p.File.Interfaces {
		if iface.Name == interfaceName {
			return iface, nil
		}
	}

	return types.Interface{}, fmt.Errorf("no interface: %s ", interfaceName)
}

// HasFunction checks whether the file has a function with the given name
func (p *baseParser) HasFunction(name string) bool {
	for _, meth := range p.File.Functions {
		if meth.Name == name {
			return true
		}
	}
	return false
}

// HasMethod checks whether the file has a method with the given name
func (p *baseParser) HasMethod(name string) bool {
	for _, meth := range p.File.Methods {
		if meth.Name == name {
			return true
		}
	}
	return false
}
