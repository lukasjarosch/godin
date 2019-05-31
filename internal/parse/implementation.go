package parse

import (
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

type implementationParser struct {
	File               *types.File
	path               string
	serviceInterface   *types.Interface
	ImplementedMethods []*types.Function
	MissingMethods     []*types.Function
}

func NewImplementationParser(path string, serviceInterface *types.Interface) *implementationParser {
	return &implementationParser{
		serviceInterface: serviceInterface,
		path:             path,
	}
}

func (i *implementationParser) Parse() (err error) {
	i.File, err = astra.ParseFile(i.path)
	if err != nil {
		return errors.Wrap(err, "Parse")

	}

	// extract all implemented and missing methods (compared to service interface)
	for _, ifaceMeth := range i.serviceInterface.Methods {
		if i.HasMethod(ifaceMeth.Name) {
			i.ImplementedMethods = append(i.ImplementedMethods, ifaceMeth)
			continue
		}
		i.MissingMethods = append(i.MissingMethods, ifaceMeth)
	}

	return nil
}

func (i *implementationParser) HasMethod(name string) bool {
	for _, meth := range i.File.Methods {
		if meth.Name == name {
			return true
		}
	}
	return false
}
