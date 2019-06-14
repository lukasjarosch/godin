package parse

import (
	"fmt"

	"strings"

	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

type transportEncodeDecodeParser struct {
	baseParser
	ImplementedMethods []string
	MissingMethods     []string
	formatStrings      []string
}

func NewTransportEncodeDecodeParser(path string, serviceIface *types.Interface) *transportEncodeDecodeParser {
	var formatStrings = []string{
		"Encode%sRequest",
		"Encode%sResponse",
		"Decode%sRequest",
		"Decode%sResponse",
	}

	return &transportEncodeDecodeParser{
		baseParser: baseParser{
			Interface: serviceIface,
			Path:      path,
		},
		formatStrings: formatStrings,
	}
}

func (p *transportEncodeDecodeParser) Parse() (err error) {
	if err := p.ParseFile(); err != nil {
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

// RequiredFunctions generates all required method names which need to exist in order for the file to be complete
func (p *transportEncodeDecodeParser) RequiredMethods() []string {
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
func (p *transportEncodeDecodeParser) EndpointName(functionName string) string {
	name := strings.Replace(functionName, "Encode", "", 1)
	name = strings.Replace(name, "Decode", "", 1)
	name = strings.Replace(name, "Request", "", 1)
	name = strings.Replace(name, "Response", "", 1)

	return name
}
