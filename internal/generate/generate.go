package generate

import (
	"os"

	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

type Generator interface {
	GenerateFull() error
	Update() error
}

type BaseGenerator struct {
	iface *types.Interface
	box   packr.Box
	opts  *Options
}

func (g *BaseGenerator) TargetExists() bool {
	if _, err := os.Stat(g.opts.TargetFile); err != nil {
		return false
	}
	return true
}

func (g *BaseGenerator) TargetPath() string {
	return g.opts.TargetFile
}

func (g *BaseGenerator) GenerateFile(box packr.Box) error {
	tpl := template.NewFile(g.opts.Template, g.opts.IsGoSource)
	data, err := tpl.Render(box, g.opts.Context)
	if err != nil {
		return errors.Wrap(err, "GenerateFile")
	}
	file := template.NewFileWriter(g.opts.TargetFile, data)
	if err := file.Write(g.opts.Overwrite); err != nil {
		return err
	}

	return nil
}

func (g *BaseGenerator) GenerateFull() error {
	if err := g.GenerateFile(g.box); err != nil {
		return errors.Wrap(err, "GenerateFull")
	}
	return nil
}
