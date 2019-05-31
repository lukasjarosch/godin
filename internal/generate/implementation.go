package generate

import (
	"os"

	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vetcher/go-astra/types"
)

type Implementation struct {
	file       string
	iface      *types.Interface
	fileExists bool
	box        packr.Box
}

func NewImplementation(box packr.Box, file string, serviceInterface *types.Interface) *Implementation {
	var exists bool
	if _, err := os.Stat(file); err != nil {
		exists = false
	} else {
		exists = true
	}

	return &Implementation{
		box: box,
		file:       file,
		iface:      serviceInterface,
		fileExists: exists,
	}
}

func (i *Implementation) GenerateFull(ctx template.Context) error {
	impl := template.NewGenerator(template.ImplementationFileOptions(ctx, i.file))
	if err := impl.GenerateFile(i.box); err != nil {
		return errors.Wrap(err, "GenerateFull")
	}
	return nil
}

func (i *Implementation) GenerateMissing(ctx template.Context) error {
	implementation := parse.NewImplementationParser(i.file, i.iface)
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

			writer := template.NewFileAppendWriter(i.file, data)
			if err := writer.Write(); err != nil {
				return errors.Wrap(err, "failed to append-write to file")
			}
			logrus.Debugf("added missing method to %s: %s", i.file, meth.String())
		}
	}

	return nil
}

func (i *Implementation) Update(ctx template.Context) error {
	if i.fileExists {
		return i.GenerateMissing(ctx)
	}
	return i.GenerateFull(ctx)
}