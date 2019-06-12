package generate

import (
	"os"

	"github.com/gobuffalo/packr"
	"github.com/vetcher/go-astra/types"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
)

type EndpointSet struct {
	file       string
	iface      *types.Interface
	fileExists bool
	box        packr.Box
}

func NewEndpointSet(box packr.Box, file string, serviceInterface *types.Interface) *EndpointSet {
	var exists bool
	if _, err := os.Stat(file); err != nil {
		exists = false
	} else {
		exists = true
	}

	return &EndpointSet{
		iface:      serviceInterface,
		file:       file,
		box:        box,
		fileExists: exists,
	}
}

func (s *EndpointSet) GenerateFull(ctx template.Context) error {
	impl := template.NewGenerator(template.FileOptions("endpoint_set", ctx, s.file))
	if err := impl.GenerateFile(s.box); err != nil {
		return errors.Wrap(err, "GenerateFull")
	}
	return nil	
}

// Update will call GenerateFull. The endpointSet cannot be updated.
func (s *EndpointSet) Update(ctx template.Context) error {
	return s.GenerateFull(ctx)
}

func (s *EndpointSet) GenerateMissing(ctx template.Context) error {
	return nil
}
