package template

import (
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

type GenerateOptions struct {
	Template   string
	IsGoSource bool
	Context    Context
	TargetFile string
	Overwrite  bool
}

type Generator struct {
	opts GenerateOptions
}

func NewGenerator(opts GenerateOptions) *Generator {
    return &Generator{
    	opts:opts,
	}
}


func (g *Generator) GenerateFile(box packr.Box) error {
	tpl := NewFile(g.opts.Template, g.opts.IsGoSource)
	data, err := tpl.Render(box, g.opts.Context)
	if err != nil {
		return errors.Wrap(err, "GenerateFile")
	}
	file := NewFileWriter(g.opts.TargetFile, data)
	if err := file.Write(g.opts.Overwrite); err != nil {
		return err
	}

	return nil
}
