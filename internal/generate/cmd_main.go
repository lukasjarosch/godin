package generate

import (
	"github.com/gobuffalo/packr"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/vetcher/go-astra/types"
	"fmt"
	"strings"
)

type CmdMain struct {
	BaseGenerator
}

func NewCmdMain(box packr.Box, serviceInterface *types.Interface, ctx template.Context, options ...Option) *CmdMain {
	defaults := &Options{
		Context:    ctx,
		Overwrite:  true,
		IsGoSource: true,
		Template:   "cmd_main_gen",
		TargetFile: fmt.Sprintf("cmd/%s/main.gen.go", strings.ToLower(ctx.Service.Name)),
	}

	for _, opt := range options {
		opt(defaults)
	}

	return &CmdMain{
		BaseGenerator{
			box:   box,
			iface: serviceInterface,
			opts:  defaults,
		},
	}
}

// Update will call GenerateFull. The cmd/main.go cannot be updated.
func (s *CmdMain) Update() error {
	return s.GenerateFull()
}

