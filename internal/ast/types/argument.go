package types

import (
	"go/ast"
)

type Argument struct {
	name, asField *ast.Ident
	typ ast.Expr
}