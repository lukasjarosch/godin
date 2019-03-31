package project

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/template"

	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/sirupsen/logrus"
)

type GodinProject struct {
	folders   []string
	templates []template.File
	Path      string
	Data      *template.Data
	Spec      *specification.Specification
}

// NewGodinProject creates an empty, preconfigured project
func NewGodinProject(spec *specification.Specification, path string) *GodinProject {

	// setup the template data with the specification
	data := &template.Data{
		ServiceName:     spec.Service.Name,
		ProjectRootPath: path,
		GrpcServiceName: spec.Service.API.Service,
		ModuleName:      spec.Project.Module,
		Spec:            spec,
	}

	return &GodinProject{
		Data: data,
		Spec: spec,
		Path: path,
	}
}

// AddFolder registers a new project folder
func (p *GodinProject) AddFolder(folder string) {
	p.folders = append(p.folders, folder)
}

// AddTemplate registers a new template to the project
func (p *GodinProject) AddTemplate(template template.File) {
	p.templates = append(p.templates, template)
}

// Render will call Render() on every registered File
func (p *GodinProject) Render() error {
	for _, tpl := range p.templates {
		if err := tpl.Render(p.Data); err != nil {
			logrus.Error(err)
			continue
		}
	}

	return nil
}

// MkdirAll creates all project folders which have been registered with AddFolder()
func (p *GodinProject) MkdirAll() error {
	for _, folder := range p.folders {
		f := p.FolderPath(folder)

		if _, err := os.Stat(f); err == nil {
			logrus.Infof("[skip] path exists %s", f)
			continue
		}

		err := os.Mkdir(f, 0755)
		if err != nil {
			return err
		}
		logrus.Infof("[mkdir] created %s", f)
	}

	return nil
}

// FolderPath returns the given (relative) path as absolute path based on the project root
func (p *GodinProject) FolderPath(subPath string) string {
	return path.Join(p.Path, subPath)
}
