package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/pkg/errors"
	"fmt"
	"github.com/spf13/viper"
)

type Endpoints struct {
	BaseGenerator
}

func NewEndpoints(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *Endpoints {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: false, // FIXME: set to 'true'
		Template:   "endpoints",
		TargetFile: "internal/service/endpoint/endpoints.go",
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &Endpoints{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

func (e *Endpoints) GenerateMissing() error {
	implementation := parse.NewEndpointsParser(e.opts.TargetFile, e.iface)
	if err := implementation.Parse(); err != nil {
		return errors.Wrap(err, "Endpoints.Parse")
	}

	if len(implementation.MissingEndpoints) > 0 {
		for _, missingEndpoint := range implementation.MissingEndpoints {

			// we miss the service name at this point since 'MethodFromType' cannot extract that
			m := template.MethodFromType(missingEndpoint)
			m.ServiceName = viper.GetString("service.name")

			tpl := template.NewPartial("endpoint", true)
			data, err := tpl.Render(e.box, m)
			if err != nil {
			    return errors.Wrap(err, "failed to render partial")
			}

			writer := template.NewFileAppendWriter(e.opts.TargetFile, data)
			if err := writer.Write(); err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to append-write to %s", e.TargetPath()))
			}
			logrus.Infof("added missing endpoint to %s: %s", e.opts.TargetFile, missingEndpoint)
		}
	}

	return nil
}

func (e *Endpoints) Update() error {
	if e.TargetExists() {
		return e.GenerateMissing()
	}
	return e.GenerateFull()
}

