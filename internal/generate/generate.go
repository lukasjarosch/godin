package generate

import (
	"os"

	"github.com/gobuffalo/packr"
	"github.com/vetcher/go-astra/types"
	"github.com/lukasjarosch/godin/internal/template"
)

type Generator interface {
	GenerateFull(ctx template.Context) error
	Update(ctx template.Context) error
}

type BaseGenerator struct {
	file  string
	iface *types.Interface
	box   packr.Box
}

func (g *BaseGenerator) TargetExists() bool {
	if _, err := os.Stat(g.file); err != nil {
		return false
	}
	return true
}


