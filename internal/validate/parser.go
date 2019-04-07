package validate

import (
	"github.com/vetcher/go-astra/types"
)

type parser struct {
	file *types.File
}

func NewParser(f *types.File) *parser {
	return &parser{
		file: f,
	}
}

func (p *parser) HasMethod(name string) bool  {
	for _, m := range p.file.Methods {
		if name == m.Name {
			return true
		}
	}
	return false
}
