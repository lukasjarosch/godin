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

func (s *Service) ParseFile() (err error) {
	s.File, err = astra.ParseFile(s.path)
	if err != nil {
		return errors.Wrap(err, "ParseFile")
	}
	return nil
}

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
	for _, meth:= range s.Interface.Methods {
		lastReturn := meth.Results[len(meth.Results) - 1]

		if lastReturn.Type.String() != "error" {
			return fmt.Errorf("last return argument must be an error: %s", meth.Name)
		}
	}

	return nil
}

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