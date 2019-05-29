package template

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal"
)

type partial struct {
	file string
	isGoSource bool
}

func NewPartial(file string, isGoSource bool) *partial {
    return &partial{
    	file:file,
    	isGoSource:isGoSource,
	}
}

func (p *partial) Render(box packr.Box, data *internal.Data) error {
	return nil	
}