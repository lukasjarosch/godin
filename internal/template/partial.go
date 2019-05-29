package template

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Partial struct {
	BaseTemplate
}

func NewPartial(name string, isGoSource bool) *Partial {
	return &Partial{
		BaseTemplate: BaseTemplate{
			name:       name,
			isGoSource: isGoSource,
		},
	}
}

func (p *Partial) Render(templateContext interface{}) ([]byte, error) {
	wd, _ := os.Getwd()

	tpl, err := template.New(p.Filename()).
		Funcs(sprig.TxtFuncMap()).
		ParseFiles(filepath.Join(wd, "templates", "partials", p.Filename()))
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	if err := tpl.Execute(&buf, templateContext); err != nil {
		return nil, err
	}

	if p.isGoSource {
		return p.FormatCode(buf.Bytes())
	}

	return buf.Bytes(), nil
}
