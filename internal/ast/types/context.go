package types

import (
	"go/ast"
)

type AstContext struct {
	Pkg *ast.Ident
	Imports []*ast.ImportSpec
	Interfaces []Iface
	Types []*ast.TypeSpec
}
