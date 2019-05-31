package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

type RequestResponse struct {
	BaseGenerator
}

func NewRequestResponse(box packr.Box, file string, serviceInterface *types.Interface) *RequestResponse {
	return &RequestResponse{
		BaseGenerator{
			box:   box,
			file:  file,
			iface: serviceInterface,
		},
	}
}

func (r *RequestResponse) Update(ctx template.Context) error {
	return r.GenerateFull(ctx)
}

func (r *RequestResponse) GenerateFull(ctx template.Context) error {
	impl := template.NewGenerator(template.FileOptions("request_response", ctx, r.file))
	if err := impl.GenerateFile(r.box); err != nil {
		return errors.Wrap(err, "GenerateFull")
	}
	return nil
}
