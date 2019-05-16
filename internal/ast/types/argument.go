package types

import (
	"go/ast"
)

type Argument struct {
	Name, asField *ast.Ident
	Typ ast.Expr
}