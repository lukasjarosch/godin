package parse

import (
	"fmt"

	"github.com/pkg/errors"
)

type amqpEncodeDecodeParser struct {
	baseParser
	publisherNames       []string
	subscriberNames      []string
	ImplementedFunctions []string
	MissingFunctions     []string
	formatStrings        []string
}

func NewAmqpEncodeDecodeParser(path string, publisherNames []string, subscriberNames []string) *amqpEncodeDecodeParser {
	return &amqpEncodeDecodeParser{
		baseParser: baseParser{
			Path: path,
		},
		publisherNames:  publisherNames,
		subscriberNames: subscriberNames,
	}
}

func (p *amqpEncodeDecodeParser) Parse() (err error) {
	if err := p.ParseFile(); err != nil {
		return errors.Wrap(err, "Parse")
	}

	// find all missing functions
	for _, function := range p.RequiredFunctions() {
		if p.HasFunction(function) {
			p.ImplementedFunctions = append(p.ImplementedFunctions, function)
			continue
		}
		p.MissingFunctions = append(p.MissingFunctions, function)
	}

	return nil
}

// RequiredEndpoints generates all required method names which need to exist in order for the file to be complete
func (p *amqpEncodeDecodeParser) RequiredFunctions() []string {
	var requiredFunctions []string

	// publishers
	for _, pub := range p.publisherNames {
		requiredFunctions = append(requiredFunctions, fmt.Sprintf("%sEncoder", pub))
	}

	return requiredFunctions
}
