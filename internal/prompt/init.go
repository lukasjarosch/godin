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
		"",
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
		"Enter the protobuf package of the service",
		"",
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


