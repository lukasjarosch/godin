package specification

import (
	"fmt"
	"strings"
)

// Service specifies a microservice
type Service struct {
	// Name is the lowercase name of the service which is used for service discovery
	Name string
	// Handler defines the name of the struct which implements the actual gRPC handler
	Handler string
	// API holds the configuration for the interface specification of the service.
	// In this case it's enforced using protocol buffers
	API Protobuf
	// Methods holds all business methods of the service including their parameters and comments
	Methods []ServiceMethod
}

// Protobuf configures the API of the service
// In future I want to extract this information directly by parsing proto files.
// For now you need to specify some information from the proto files
type Protobuf struct {
	// Service is the name of the service which is being implemented by this microservice
	Service string
	// Package holds the full package of the protobuf file
	Package string
}

// ServiceMethod defines the base data structure for representing a service method
// One ServiceMethod represents a RPC defined by the API.
type ServiceMethod struct {
	// Name defines the methods name
	Name string
	// Comments holds the methods's comment. One element per comment-line
	Comments []string
	// Arguments holds all input arguments to the method (the data-types)
	Arguments []Variable
	// Returns holds all output data-types
	Returns []Variable
}

// Signature returns the methods's signature as string.
// No checks are performed at this stage!
func (m ServiceMethod) Signature() string {
	format := "func %s(%s) (%s)"

	var args []string
	for _, arg := range m.Arguments {
		if arg.String() != "" {
			args = append(args, arg.String())
		}
	}

	var returns []string
	for _, ret := range m.Returns {
		returns = append(returns, ret.String())
	}

	return fmt.Sprintf(format, m.Name, strings.Join(args, ", "), strings.Join(returns, ", "))
}

// Variable
type Variable struct {
	Type string
	Name string
}

func (a Variable) String() string {
	if a.Name == "" {
		return a.Type
	}
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}
