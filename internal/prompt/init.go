package prompt

import (
	"os"
)

func ServiceName() string {
	prompt := NewPrompt(
		"Enter the service name (lowercase)",
		"",
		Validate(
			MinLengthThree(),
			Lowercase(),
		),
	)
	serviceName, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return serviceName
}

func ServiceNamespace() string {
	prompt := NewPrompt(
		"Enter the namespace (lowercase)",
		"",
		Validate(
			MinLengthThree(),
			Lowercase(),
		),
	)
	namespace, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return namespace
}

func ServiceModule() string {
	prompt := NewPrompt(
		"Enter the Go module name",
		"github.com/username/service",
		Validate(
			MinLengthThree(),
		),
	)
	module, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return module
}

func ProtoServiceName() string {
	prompt := NewPrompt(
		"Enter the gRPC service name which this service implements",
		"",
		Validate(
			MinLengthThree(),
		),
	)
	protoService, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return protoService
}

func ProtoPackage() string {
	prompt := NewPrompt(
		"Enter the importable go-package of the protobuf service stubs (protoc go_out)",
		"github.com/username/protobuf-go",
		Validate(
			MinLengthThree(),
		),
	)
	protoPackage, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return protoPackage
}

func ProtoPath() string {
	prompt := NewPrompt(
		"Absolute path to the .proto file which defines the service",
		"/home/lukas/devel/work/protobuf",
		Validate(
			MinLengthThree(),
			ProtoFileExtension(),
		),
	)
	protoPath, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return protoPath
}

func DockerRegistry() string {
	prompt := NewPrompt(
		"Enter your docker registry",
		"registry.hub.docker.com",
		Validate(
			MinLengthThree(),
		),
	)
	registry, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	return registry
}
