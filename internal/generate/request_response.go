package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
)

type RequestResponse struct {
	BaseGenerator
}

func NewRequestResponse(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *RequestResponse {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "request_response",
		TargetFile: "internal/service/endpoint/request_response.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &RequestResponse{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update is disabled for this file, it will only proxy the call to GenerateFull()
func (r *RequestResponse) Update() error {
	return r.GenerateFull()
}

