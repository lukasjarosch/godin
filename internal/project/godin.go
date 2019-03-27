package project

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/template"

	"github.com/sirupsen/logrus"
)

type GodinProject struct {
	serviceName string
	path        string // absolute path to the root of the godin project
	folders     []string
	templates   []template.File
}

// NewGodinProject creates an empty, preconfigured project
func NewGodinProject(serviceName string, path string) *GodinProject {
	logrus.Infof("creating godin project '%s' in %s", serviceName, path)

	return &GodinProject{
		serviceName: serviceName,
		path:        path,
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
		if err := tpl.Render(); err != nil {
			logrus.Error(err)
			continue
		}
	}

	return nil
}

// MkdirAll creates all project folders which have been registered with AddFolder()
func (p *GodinProject) MkdirAll() error {
	for _, folder := range p.folders {
		f := p.Path(folder)

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

// Path returns the given (relative) path as absolute path based on the project root
func (p *GodinProject) Path(subPath string) string {
	return path.Join(p.path, subPath)
}
