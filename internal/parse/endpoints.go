package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

type endpointsParser struct {
	baseParser
	ExistingEndpoints []*types.Function
	MissingEndpoints  []*types.Function
}

func NewEndpointsParser(path string, serviceIface *types.Interface) *endpointsParser {
	return &endpointsParser{
		baseParser: baseParser{
			Interface: serviceIface,
			Path:      path,
		},
	}
}

func (p *endpointsParser) Parse() (err error) {
	if err := p.ParseFile(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	// find all missing endpoints
	for _, endpoint := range p.Interface.Methods {
		if p.HasFunction(fmt.Sprintf("%sEndpoint", endpoint.Name)) {
			p.ExistingEndpoints = append(p.ExistingEndpoints, endpoint)
			continue
		}
		p.MissingEndpoints = append(p.MissingEndpoints, endpoint)
	}

	return nil
}
