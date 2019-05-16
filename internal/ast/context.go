package ast

import (
	"go/ast"
	"go/token"

	"fmt"

	"github.com/lukasjarosch/godin/internal/ast/types"
)

type ServiceFileContext struct {
	File       *ast.File
	FSet       *token.FileSet
	Pkg        *ast.Ident
	Imports    []*ast.ImportSpec
	Interfaces []types.Iface
	Types      []*ast.TypeSpec
}

func (ctx ServiceFileContext) Validate() error {
	if len(ctx.Interfaces) != 1 {
		return fmt.Errorf("service file must contain exactly 1 interface, found %d", len(ctx.Interfaces))
	}
	return nil
}
