package parse

import (
	"fmt"

	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

type transportRequestResponseParser struct {
	baseParser
	ImplementedFunctions []string
	MissingFunctions     []string
	formatStrings        []string
}

func NewTransportRequestResponseParser(path string, serviceIface *types.Interface) *transportRequestResponseParser {
	var formatStrings = []string{
		"Encode%sRequest",
		"Encode%sResponse",
		"Decode%sRequest",
		"Decode%sResponse",
	}

	return &transportRequestResponseParser{
		baseParser: baseParser{
			Interface: serviceIface,
			Path:      path,
		},
		formatStrings: formatStrings,
	}
}

func (p *transportRequestResponseParser) Parse() (err error) {
	if err := p.ParseFile(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	// find all missing functions
	for _, meth := range p.RequiredFunctions() {
		if p.HasFunction(meth) {
			p.ImplementedFunctions = append(p.ImplementedFunctions, meth)
			continue
		}
		p.MissingFunctions = append(p.MissingFunctions, meth)
	}

	return nil
}

// RequiredFunctions generates all required method names which need to exist in order for the file to be complete
func (p *transportRequestResponseParser) RequiredFunctions() []string {
	var requiredMethods []string

	for _, meth := range p.Interface.Methods {
		for n := 0; n <= len(p.formatStrings)-1; n++ {
			method := fmt.Sprintf(p.formatStrings[n], meth.Name)
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

