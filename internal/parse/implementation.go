package parse

import (
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

type implementationParser struct {
	baseParser
	ImplementedMethods []*types.Function
	MissingMethods     []*types.Function
}

func NewImplementationParser(path string, serviceInterface *types.Interface) *implementationParser {
	return &implementationParser{
		baseParser: baseParser{
			Interface: serviceInterface,
			Path:      path,
		},
	}
}

func (i *implementationParser) Parse() (err error) {
	i.File, err = astra.ParseFile(i.Path)
	if err != nil {
		return errors.Wrap(err, "Parse")

	}

	// extract all implemented and missing methods (compared to service interface)
	for _, ifaceMeth := range i.Interface.Methods {
		if i.HasMethod(ifaceMeth.Name) {
			i.ImplementedMethods = append(i.ImplementedMethods, ifaceMeth)
			continue
		}
		i.MissingMethods = append(i.MissingMethods, ifaceMeth)
	}

	return nil
}

