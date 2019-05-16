package template

import (
	"bytes"
	"go/format"
	"os"
	tpl "text/template"

	"github.com/lukasjarosch/godin/internal"
	"github.com/Masterminds/sprig"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
)

// File defines the interface for our template-files
type File interface {
	Render(box packr.Box, data *internal.Data) error
}

// file defines a single template-file
type file struct {
	Name       string // Name is the template filename
	TargetPath string // Absolute path, output path
	goSource   bool   // Set and the template will be formatted
}

// NewTemplateFile returns a new file
func NewTemplateFile(name string, path string, goSource bool) *file {
	return &file{
		Name:       name,
		TargetPath: path,
		goSource:   goSource,
	}
}

// Render will parse the template and write it out as file
// If the template is defined as goSource, the renderGoCode() function is called for processing
// Every other template is written using template.Execute()
//
// TODO: Catch file exists errors and handle them, better not overwrite things :)
func (t *file) Render(box packr.Box, data *internal.Data) error {

	stat, _ := os.Stat(t.TargetPath)
	if stat != nil {
		logrus.Infof("[skip] template already written: %s", t.Name)
		return nil
	}

	templateData, err := box.FindString(t.Name)
	if err != nil {
		return err
	}

	template, err := tpl.New(t.TargetPath).Funcs(sprig.TxtFuncMap()).Parse(templateData)
	f, err := os.Create(t.TargetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if t.goSource {
		err := t.renderGoCode(f, template, data)
		if err != nil {
			return err
		}
		logrus.Infof("[template] rendered %s", t.Name)
		return nil
	}

	err = template.Execute(f, data)
	if err != nil {
		return err
	}
	logrus.Infof("[template] rendered %s", t.Name)

	return nil
}

// renderGoCode parses the given file using the given template. The parsed file is
// written into a bytes.Buffer which is used to format the source before writing the file.
func (t *file) renderGoCode(f *os.File, template *tpl.Template, data *internal.Data) error {
	var out bytes.Buffer

	err := template.Execute(&out, data)
	if err != nil {
		return err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		logrus.Info(template.ParseName)
		return err
	}

	_, err = f.Write(formatted)
	if err != nil {
		return err
	}

	return nil
}
