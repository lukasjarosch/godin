package template

import (
	config "github.com/spf13/viper"
)

var fileOptions = map[string]GenerateOptions{
	"service_stub": {
		Template:   "service_stub",
		IsGoSource: true,
		TargetFile: "internal/service/service.go",
		Overwrite:  true,
	},
	"logging_middleware": {
		Template:   "logging_middleware",
		IsGoSource: true,
		TargetFile: "internal/service/middleware/logging.go",
		Overwrite:  true,
	},
	"middleware": {
		Template:   "middleware",
		IsGoSource: true,
		TargetFile: "internal/service/middleware/middleware.go",
		Overwrite:  true,
	},
	"request_response": {
		Template:   "request_response",
		IsGoSource: true,
		TargetFile: "internal/service/endpoint/request_response.go",
		Overwrite:  true,
	},
	"endpoint_set": {
		Template:   "endpoint_set",
		IsGoSource: true,
		TargetFile: "internal/service/endpoint/set.go",
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
	"implementation": {
		Overwrite:  true,
		IsGoSource: true,
		Template:   "implementation",
	},
}

func FileOptions(name string, tplContext Context, targetPath string) GenerateOptions {
	ctx := fileOptions[name]
	ctx.TargetFile = targetPath
	ctx.Context = tplContext

	return ctx
}

func ImplementationFileOptions(tplContext Context, targetPath string) GenerateOptions {
	ctx := fileOptions["implementation"]
	ctx.TargetFile = targetPath
	ctx.Context = tplContext

	return ctx
}

func MiddlewareOptions() GenerateOptions {
	ctx := Context{
		Service: Service{
			Name:   config.GetString("service.name"),
			Module: config.GetString("service.module"),
		},
	}

	opts := fileOptions["middleware"]
	opts.Context = ctx

	return opts
}

func RequestResponseOptions(ctx Context) GenerateOptions {
	opts := fileOptions["request_response"]
	opts.Context = ctx

	return opts
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
