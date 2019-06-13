package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
	"strings"
)

type transportRequestResponseParser struct {
	File               *types.File
	path               string
	serviceInterface   *types.Interface
	ImplementedMethods []string
	MissingMethods     []string
}

var formatStrings = []string{
	"Encode%sRequest",
	"Encode%sResponse",
	"Decode%sRequest",
	"Decode%sResponse",
}

func NewTransportRequestResponseParser(path string, serviceIface *types.Interface) *transportRequestResponseParser {
	return &transportRequestResponseParser{
		serviceInterface: serviceIface,
		path:             path,
	}
}

func (p *transportRequestResponseParser) Parse() (err error) {
	p.File, err = astra.ParseFile(p.path)
	if err != nil {
		return errors.Wrap(err, "Parse")
	}

	// find all missing methods
	for _, meth := range p.RequiredMethods() {
		if p.HasMethod(meth) {
			p.ImplementedMethods = append(p.ImplementedMethods, meth)
			continue
		}
		p.MissingMethods = append(p.MissingMethods, meth)
	}

	return nil
}

// RequiredMethods generates all required method names which need to exist in order for the file to be complete
func (p *transportRequestResponseParser) RequiredMethods() []string {
	var requiredMethods []string

	for _, meth := range p.serviceInterface.Methods {
		for n := 0; n <= len(formatStrings)-1; n++ {
			method := fmt.Sprintf(formatStrings[n], meth.Name)
			requiredMethods = append(requiredMethods, method)
		}
	}

	return requiredMethods
}

// EndpointName will extract the endpoint name from a function.
// For example: EncodeHelloRequest => Hello
func (p *transportRequestResponseParser) EndpointName(functionName string) string {
	name := strings.Replace(functionName, "Encode", "", 1)
	name = strings.Replace(name, "Decode", "", 1)
	name = strings.Replace(name, "Request", "", 1)
	name = strings.Replace(name, "Response", "", 1)

	return name
}

func (p *transportRequestResponseParser) HasMethod(name string) bool {
	for _, meth := range p.File.Functions {
		if meth.Name == name {
			return true
		}
	}
	return false
}
