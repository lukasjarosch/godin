package template

import (
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	config "github.com/spf13/viper"
)

func WriteDockerfile(box packr.Box) error {
	dockerfile := NewFile("Dockerfile", false)
	data, err := dockerfile.Render(box, Context{
		Service: Service{
			Name: config.GetString("service.name"),
		},
	})
	if err != nil {
		return errors.Wrap(err, "CreateDockerfile")
	}
	file := NewFileWriter("Dockerfile", data)
	if err := file.Write(true); err != nil {
		return err
	}

	return nil
}
