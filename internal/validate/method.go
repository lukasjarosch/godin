package validate

import (
	"errors"

	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/vetcher/go-astra/types"

	"fmt"
)

type method struct {
	spec *specification.ServiceMethod
	impl *types.Method
}

func ValidateMethod(receiver string, spec *specification.ServiceMethod, impl *types.Method) error {

	meth := &method{spec:spec, impl:impl}

	if len(spec.Arguments) != len(impl.Args) {
		return errors.New(fmt.Sprintf("ARGUMENT-MISMATCH:\n\tspec: %s\n\timpl: %s", spec.Signature(), impl.Function.String()))
	}
	if len(impl.Results) != len(spec.Returns) {
		return errors.New(fmt.Sprintf("RESULT-MISMATCH:\n\tspec: %s\n\timpl: %s", spec.Signature(), impl.Function.String()))
	}

	// validate param args
	if err := meth.validateArgs(spec.Arguments, impl.Args); err != nil {
		return err
	}

	// validate return args
	if err := meth.validateArgs(spec.Returns, impl.Results); err != nil {
		return err
	}

	return nil
}

func (m *method) validateArgs(spec []specification.Variable, impl []types.Variable) error {

	argChan := make(chan specification.Variable, len(spec))
	for _, specArg := range spec {
		argChan <- specArg
	}
	close(argChan)

	implChan := make(chan types.Variable, len(impl))
	for _, implArg := range impl{
		implChan <- implArg
	}
	close(implChan)

	for specArg := range argChan {
		for implArg := range implChan {
			if specArg.Name == implArg.Name &&
				specArg.Type == implArg.Type.String() {
				break
			}
			return errors.New(fmt.Sprintf("ARGUMENT-MISMATCH:\n\tspec: %s\n\timpl: %s", m.spec.Signature(), m.impl.Function.String()))
		}
	}

	return nil
}