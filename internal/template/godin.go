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
	"k8s_service": {
		Template:   "k8s_service",
		IsGoSource: false,
		TargetFile: "k8s/service.yaml",
		Overwrite:  true,
	},
	"k8s_deployment": {
		Template:   "k8s_deployment",
		IsGoSource: false,
		TargetFile: "k8s/deployment.yaml",
		Overwrite:  true,
	},
	"makefile": {
		Template:   "makefile",
		IsGoSource: false,
		TargetFile: "Makefile",
		Overwrite:  false,
	},
	"errors": {
		Template:   "domain_errors",
		IsGoSource: true,
		TargetFile: "internal/service/domain/errors.go",
		Overwrite:  false,
	},
}

func FileOptions(name string, tplContext Context, targetPath string) GenerateOptions {
	ctx := fileOptions[name]
	ctx.TargetFile = targetPath
	ctx.Context = tplContext

	return ctx
}

func K8sDeploymentOptions() GenerateOptions {
	ctx := Context{
		Service: Service{
			Name:      config.GetString("service.name"),
			Namespace: config.GetString("service.namespace"),
		},
		Docker: Docker{
			Registry: config.GetString("docker.registry"),
		},
	}
	opts := fileOptions["k8s_deployment"]
	opts.Context = ctx

	return opts
}

func K8sServiceOptions() GenerateOptions {
	ctx := Context{
		Service: Service{
			Name:      config.GetString("service.name"),
			Namespace: config.GetString("service.namespace"),
		},
	}
	opts := fileOptions["k8s_service"]
	opts.Context = ctx

	return opts
}

func MakefileOptions(ctx Context) GenerateOptions {
	opts := fileOptions["makefile"]
	opts.Context = ctx

	return opts
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

func DomainErrorsOptions() GenerateOptions {
	ctx := Context{}

	opts := fileOptions["errors"]
	opts.Context = ctx

	return opts
}
