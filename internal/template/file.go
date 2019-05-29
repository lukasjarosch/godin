package template

import (
	"bytes"
	"path"
	"text/template"

	"fmt"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

// file template rendering
// a template-file can include one/more partials
// each file therefore needs a different set of data

type File struct {
	BaseTemplate
	tpl       *template.Template
	templates []string
}

func NewFile(name string, isGoSource bool) *File {
	wd, _ := os.Getwd()

	return &File{
		BaseTemplate: BaseTemplate{
			name:       name,
			isGoSource: isGoSource,
			rootPath:   wd,
		},
	}
}

// prepare the templates paths which might be included by the loaded template file
func (f *File) prepare(fs packr.Box) error {
	f.templates = append(f.templates, f.TemplatePath())
	f.templates = append(f.templates, "partials/service_method.tpl")

	return nil
}

// Render the specified template file
func (f *File) Render(fs packr.Box, templateContext Context) (rendered []byte, err error) {
	if err := f.prepare(fs); err != nil {
		return nil, err
	}

	var templateData string
	for _, tpl := range f.templates {
		tmp, err := fs.FindString(tpl)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("FindString: %s", tpl))
		}
		templateData += tmp
	}

	buf := bytes.Buffer{}
	f.tpl, err = template.New(path.Base(f.templates[0])).Funcs(sprig.TxtFuncMap()).Parse(templateData)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}


	if err := f.tpl.Execute(&buf, templateContext); err != nil {
		return nil, errors.Wrap(err, "Execute")
	}

	if f.isGoSource {
		return f.FormatCode(buf.Bytes())
	}

	return buf.Bytes(), nil
}
