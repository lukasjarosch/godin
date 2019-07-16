package generate

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vetcher/go-astra/types"
	"path/filepath"
	config "github.com/spf13/viper"
)

type Implementation struct {
	BaseGenerator
}

func NewImplementation(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *Implementation {
	defaults := &Options{
		Context:ctx,
		Overwrite:true,
		IsGoSource:true,
		Template:"implementation",
		TargetFile: filepath.Join("internal", "service", "usecase", fmt.Sprintf("%s.go", config.GetString("service.name"))),
	}

	for _, opt := range options {
		opt(defaults)
	}


	return &Implementation{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:defaults,
		},
	}
}

func (i *Implementation) GenerateMissing() error {
	implementation := parse.NewImplementationParser(i.opts.TargetFile, i.iface)
	if err := implementation.Parse(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	if len(implementation.MissingMethods) > 0 {
		for _, meth := range implementation.MissingMethods {
			tpl := template.NewPartial("service_method", true)
			data, err := tpl.Render(i.box, template.MethodFromType(meth))
			if err != nil {
				return errors.Wrap(err, "failed to render partial")
			}

			writer := template.NewFileAppendWriter(i.opts.TargetFile, data)
			if err := writer.Write(); err != nil {
				return errors.Wrap(err, "failed to append-write to file")
			}
			logrus.Infof("added missing method to %s: %s", i.opts.TargetFile, meth.String())
		}
	}

	return nil
}

func (i *Implementation) Update() error {
	if i.TargetExists() {
		return i.GenerateMissing()
	}
	return i.GenerateFull()
}
