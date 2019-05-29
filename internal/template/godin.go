package template

import (
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	config "github.com/spf13/viper"
)

// WriteDockerfile will render the Dockerfile.tpl template to ./Dockerfile, overwriting existing files
func WriteDockerfile(box packr.Box) error {
	dockerfile := NewFile("Dockerfile", false)
	data, err := dockerfile.Render(box, Context{
		Service: Service{
			Name: config.GetString("service.name"),
		},
	})
	if err != nil {
		return errors.Wrap(err, "WriteDockerfile")
	}
	file := NewFileWriter("Dockerfile", data)
	if err := file.Write(true); err != nil {
		return err
	}

	return nil
}

func WriteGitignore(box packr.Box) error {
	gitignore := NewFile("gitignore", false)
	data, err := gitignore.Render(box, Context{})
	if err != nil {
		return errors.Wrap(err, "WriteGitignore")
	}
	file := NewFileWriter(".gitignore", data)
	if err := file.Write(true); err != nil {
		return err
	}

	return nil
}