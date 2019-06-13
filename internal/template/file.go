package template

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

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
	f.templates = append(f.templates, "partials/logging_method.tpl")
	f.templates = append(f.templates, "partials/request.tpl")
	f.templates = append(f.templates, "partials/response.tpl")
	f.templates = append(f.templates, "partials/grpc_encode_request.tpl")
	f.templates = append(f.templates, "partials/grpc_decode_request.tpl")
	f.templates = append(f.templates, "partials/grpc_encode_response.tpl")
	f.templates = append(f.templates, "partials/grpc_decode_response.tpl")

	return nil
}

// Render the specified template file
func (f *File) Render(fs packr.Box, templateContext Context) (rendered []byte, err error) {
	if err := f.prepare(fs); err != nil {
		return nil, err
	}

	// parse base template first
	tmp, err := fs.FindString(f.templates[0])
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("FindString: %s", f.templates[0]))
	}
	f.tpl, err = template.New(path.Base(f.templates[0])).Funcs(sprig.TxtFuncMap()).Parse(tmp)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}

	// parse all loaded partial templates
	for i := 1; i < len(f.templates); i++ {
		tmp, err := fs.FindString(f.templates[i])
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("FindString: %s", f.templates[i]))
		}
		f.tpl, err = f.tpl.Parse(tmp)
		if err != nil {
			return nil, errors.Wrap(err, "Parse")
		}
	}

	buf := bytes.Buffer{}
	if err := f.tpl.Execute(&buf, templateContext); err != nil {
		return nil, errors.Wrap(err, "Execute")
	}

	if f.isGoSource {
		return f.FormatCode(buf.Bytes())
	}

	return buf.Bytes(), nil
}
