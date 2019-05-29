package parse

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	. "github.com/vetcher/go-astra/types"
)

type Service struct {
	path      string
	File      *File
	Interface *Interface
}

func NewServiceParser(path string) *Service {
	return &Service{
		path: path,
	}
}

// ParseFile will use go-astra to parse the service-file
func (s *Service) ParseFile() (err error) {
	s.File, err = astra.ParseFile(s.path)
	if err != nil {
		return errors.Wrap(err, "ParseFile")
	}
	return nil
}

// FindInterface searches for an interface with the given name. If the interface is found, it's set to the Interface field
func (s *Service) FindInterface(interfaceName string) error {
	for _, iface := range s.File.Interfaces {
		if iface.Name == interfaceName {
			s.Interface = &iface
			return nil
		}
	}

	return fmt.Errorf("no interface: %s ", interfaceName)
}

// ValidateInterface will check if the interface meets Godin's requirements:
//
// + every method must have at least one comment line
// + first parameter must be 'context.Context'
// + last return parameter must be ' error'
// + all results must be named
// + if a custom type is used, it MUST be defined in the same file
func (s *Service) ValidateInterface() (err error) {
	err = s.validateMethodComments()
	if err != nil {
		return err
	}

	err = s.validateContextParameter()
	if err != nil {
		return err
	}

	err = s.validateErrorReturn()
	if err != nil {
		return err
	}

	err = s.validateNamedResults()
	if err != nil {
		return err
	}

	err = s.validateCustomTypes()
	if err != nil {
		return err
	}

	return nil
}

// validateMethodComments iterates over all interface methods and checks whether each method has at least one comment line.
// If no comment is found, the function returns with an error
func (s *Service) validateMethodComments() error {
	for _, meth := range s.Interface.Methods {
		if len(meth.Docs) < 1 {
			return fmt.Errorf("%s endpoint is missing a comment", meth.Name)
		}
	}
	return nil
}

// validateContextParameter iterates over all interface methods and checks whether each first argument is a context.
// An error is returned if a method does not have the context as first argument.
func (s *Service) validateContextParameter() error {
	for _, meth := range s.Interface.Methods {
		typ := strings.ToLower(meth.Args[0].Type.String())
		if !strings.Contains(typ, "context") {
			return fmt.Errorf("first argument must be a context: %s", meth.Name)
		}
	}

	return nil
}

// validateErrorReturn iterates over all interface methods and checks whether every last argument is an error.
// An error is returned if a method does not have an error as last return argument.
func (s *Service) validateErrorReturn() error {
	for _, meth := range s.Interface.Methods {
		lastReturn := meth.Results[len(meth.Results)-1]

		if lastReturn.Type.String() != "error" {
			return fmt.Errorf("last return argument must be an error: %s", meth.Name)
		}
	}

	return nil
}

// validateNamedResults ensures that all return parameters are named. An error is returned if an unnamed param is found.
func (s *Service) validateNamedResults() error {
	for _, meth := range s.Interface.Methods {
		for _, arg := range meth.Results {
			if arg.Name == "" {
				return fmt.Errorf("results must be named: %s", meth.Name)
			}
		}
	}
	return nil
}

// validateCustomTypes will try and search for any custom types inside the service.go file
// if the type is not declared in the same file, an error is returned.
func (s *Service) validateCustomTypes() (err error) {
	for _, meth := range s.Interface.Methods {
		for _, arg := range meth.Args {
			if !IsBuiltin(arg.Type) {
				err = s.findCustomTypeDeclaration(arg.Type.String())
			}
		}

		for _, res := range meth.Results {
			if !IsBuiltin(res.Type) {
				err = s.findCustomTypeDeclaration(res.Type.String())
			}
		}
	}
	return err
}

// findCustomTypeDeclaration searches for a given typeName inside the current file.
// It will search in structs and types. It will also try to search in the imports, but one should
// not rely on that to work. It's just present to ensure that we can use context.Context in the service-file.
func (s *Service) findCustomTypeDeclaration(name string) error {
	for _, s := range s.File.Structures {
		if s.Name == name {
			return nil
		}
	}

	for _, t := range s.File.Types {
		if t.Name == name {
			return nil
		}
	}

	// maybe the type is imported (e.g. context.Context)
	// this is not 100% acurate, but i should get the job done
	for _, i := range s.File.Imports {
		if strings.Contains(strings.ToLower(name), i.Package) {
			return nil
		}
	}

	return fmt.Errorf("type %s is not defined in service.go", name)
}
