package specification

import (
	"fmt"
	"strings"
	"path"
)

// Service specifies a microservice
type Service struct {
	// Name is the lowercase name of the service which is used for service discovery
	Name string
	// Description of the service
	Description string
	// Handler defines the name of the struct which implements the actual gRPC handler
	Handler string
	// API holds the configuration for the interface specification of the service.
	// In this case it's enforced using protocol buffers
	API Protobuf
	// Methods holds all business methods of the service including their parameters and comments
	Methods []ServiceMethod
	// All service dependencies
	Dependencies []Dependency
	// Errors specifies all errors of the service's business domain
	Errors []ErrorSpec
}

// HasMethod checks whether a method with a given name exists in the service specification
func (s *Service) HasMethod(name string) bool {
	for _, m := range s.Methods {
		if name == m.Name {
			return true
		}
	}
	return false
}

func (s *Service) GetMethod(name string) *ServiceMethod {
	for _, m := range s.Methods {
		if name == m.Name {
			return &m
		}
	}
	return nil
}

// Protobuf configures the API of the service
// In future I want to extract this information directly by parsing proto files.
// For now you need to specify some information from the proto files
type Protobuf struct {
	// Service is the name of the service which is being implemented by this microservice
	Service string
	// Package holds the full package of the protobuf file
	Package string
	// Modules stores the go-module in which the protobuf stubs live
	Module string
}

// Import returns a full importable module string for the protobufs
// This will only work if your proto packages are 1:1 mappable to your folder structure of the protobuf-specs
func (p Protobuf) Import() string {
	packagePath := strings.Replace(p.Package, ".", "/", -1)
	return path.Join(p.Module, packagePath)
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

func (v Variable) DefaultValue(specification *Specification) string {
	switch v.Type {
	case "string":
		return ""
	case "int64":
	case "int32":
	case "int":
		return "0"
	case "boolean":
		return "false"
	case "error":
		return "nil"
	}

	for _, s := range specification.Models.Structs {
		if strings.ToLower(v.Name) == strings.ToLower(s.Name) {
			return fmt.Sprintf("%s{}", s.Name)
		}
	}

	for _, e := range specification.Models.Enums {
		if strings.ToLower(e.Name) == strings.ToLower(v.Name) {
			return "0"
		}
	}

	return ErrModelUnspecified.Error()
}

func (a Variable) String() string {
	if a.Name == "" {
		return a.Type
	}
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}
