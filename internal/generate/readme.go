package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type Readme struct {
	BaseGenerator
}

func NewReadme(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *Readme {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: false,
		Template:   "readme",
		TargetFile: "README.md",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &Readme{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

func (s *Readme) Update() error {
	return s.GenerateFull()
}
