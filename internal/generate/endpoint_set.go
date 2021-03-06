package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type EndpointSet struct {
	BaseGenerator
}

func NewEndpointSet(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *EndpointSet {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "endpoint_set",
		TargetFile: "internal/service/endpoint/set.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &EndpointSet{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull. The endpointSet cannot be updated.
func (s *EndpointSet) Update() error {
	return s.GenerateFull()
}

