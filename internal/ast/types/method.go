package types

import (
	"go/ast"
)

type Method struct {
	Name            *ast.Ident
	Params          []Argument
	Results         []Argument
}
