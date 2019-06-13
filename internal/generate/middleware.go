package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type Middleware struct {
	BaseGenerator
}

func NewMiddleware(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *Middleware {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "middleware",
		TargetFile: "internal/service/middleware/middleware.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &Middleware{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

func (m *Middleware) Update() error {
	return m.GenerateFull()
}
