package project

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/template"

	"os/exec"

	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GodinProject struct {
	folders   []string
	templates []template.File
	Path      string
	Data      *internal.Data
	Spec      *specification.Specification
	box       packr.Box
}

const ConfigFile = "godin.toml"

// HasConfig checks whether the config is loadable.
// If it is, it will automatically be loaded and true is returned.
func HasConfig() error {
	cwd, _ := os.Getwd()
	viper.SetConfigName("godin")
	viper.SetConfigType("toml")
	viper.AddConfigPath(cwd)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func SaveConfig() {
	viper.WriteConfigAs(ConfigFile)
}

// NewGodinProject creates an empty, preconfigured project
func NewGodinProject(data *internal.Data, box packr.Box) *GodinProject {
	return &GodinProject{
		Data: data,
		box:  box,
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
		if err := tpl.Render(p.box, p.Data); err != nil {
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

// InitModule will initialize the module with the given name in the current directory
func (p *GodinProject) InitModule(name string) error {
	modCmd := exec.Command("go", "mod", "init", name)
	err := modCmd.Run()
	if err != nil {
		return err
	}

	return nil
}
