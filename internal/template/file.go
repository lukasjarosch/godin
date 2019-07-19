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
	config "github.com/spf13/viper"
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

// prepare loads all template paths, starting with the requested base-template
func (f *File) prepare(fs packr.Box) error {
	f.templates = append(f.templates, f.TemplatePath())

	for _, partial := range PartialTemplates {
		f.templates = append(f.templates, partial)
	}

	return nil
}

// Render the specified template file
func (f *File) Render(fs packr.Box, templateContext Context) (rendered []byte, err error) {
	skipPartials := false
	if err := f.prepare(fs); err != nil {
		return nil, err
	}

	// parse base template first
	tmp, err := fs.FindString(f.templates[0])
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("FindString: %s", f.templates[0]))
	}
	f.tpl = template.New(path.Base(f.templates[0])).Funcs(sprig.TxtFuncMap()).Funcs(map[string]interface{}{
		"ReadmeOptionCheckbox": func(option string) string {
			if !config.IsSet(option) {
				img := "![disabled](https://img.icons8.com/color/24/000000/close-window.png)"
				return img
			}

			if config.GetBool(option) {
				img := "![enabled](https://img.icons8.com/color/24/000000/checked.png)"
				return img
			}
			img := "![disabled](https://img.icons8.com/color/24/000000/close-window.png)"
			return img

		},
	})

	// FIXME: hack to support Makefiles
	if f.name == "makefile" {
		f.tpl.Delims("<<", ">>")
		skipPartials = true
	}

	f.tpl, err = f.tpl.Parse(tmp)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}

	// parse all loaded partial templates
	if !skipPartials {
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
