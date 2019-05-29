package template

import (
	config "github.com/spf13/viper"
)

var fileOptions = map[string]GenerateOptions{
	"service_stub": {
		Template:   "service_stub",
		IsGoSource: false,
		TargetFile: "internal/service/service.go",
		Overwrite:  true,
	},
	"dockerfile": {
		Template:   "Dockerfile",
		IsGoSource: false,
		TargetFile: "Dockerfile",
		Overwrite:  true,
	},
	"gitignore": {
		Template:   "gitignore",
		IsGoSource: false,
		TargetFile: ".gitignore",
		Overwrite:  true,
	},
}

func DockerfileOptions() GenerateOptions {
	ctx := Context{
		Service: Service{
			Name: config.GetString("service.name"),
		},
	}

	opts := fileOptions["dockerfile"]
	opts.Context = ctx

	return opts
}

func ServiceStubOptions() GenerateOptions {
	ctx := Context{
		Service: Service{
			Name: config.GetString("service.name"),
		},
	}

	opts := fileOptions["service_stub"]
	opts.Context = ctx

	return opts
}

func GitignoreOptions() GenerateOptions {
	ctx := Context{}

	opts := fileOptions["gitignore"]
	opts.Context = ctx

	return opts
}
