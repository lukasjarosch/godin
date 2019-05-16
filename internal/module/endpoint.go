package module

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/ast"
	"github.com/lukasjarosch/godin/internal/template"
)

type Endpoint struct {
	Data *template.Data
}

func NewEndpoint(data *template.Data) *Endpoint {
	return &Endpoint{Data: data}
}

func (e *Endpoint) Execute() error {

	f, err := os.Open(path.Join(e.Data.Project.RootPath, "internal", "service.go"))
	if err != nil {
		return err
	}

	parser := ast.NewFile("service.go", f)
	if err := parser.Process(); err != nil {
		return err
	}

	return nil
}

func (e *Endpoint) prompt() (string, error) {
	return "", nil
}
