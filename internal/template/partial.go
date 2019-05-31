package template

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	"fmt"
)

type Partial struct {
	BaseTemplate
	tpl *template.Template
}

func NewPartial(name string, isGoSource bool) *Partial {
	return &Partial{
		BaseTemplate: BaseTemplate{
			name:       name,
			isGoSource: isGoSource,
		},
	}
}

func (p *Partial) Render(fs packr.Box, templateContext interface{}) (rendered []byte, err error) {

	templateData, err := fs.FindString(filepath.Join("partials", p.Filename()))
	if err != nil {
		return nil, errors.Wrap(err, "FindString")
	}

	// in order to use the partial template definitions, we need to include them somewhere
	// thus we just add the template to use itself
	templateData += fmt.Sprintf("{{ template \"%s\" . }}", p.name)

	p.tpl, err = template.New(p.Filename()).Funcs(sprig.TxtFuncMap()).Parse(templateData)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}

	buf := bytes.Buffer{}
	if err := p.tpl.Execute(&buf, templateContext); err != nil {
		return nil, err
	}

	if p.isGoSource {
		return p.FormatCode(buf.Bytes())
	}

	return buf.Bytes(), nil
}
