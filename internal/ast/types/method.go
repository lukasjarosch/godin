package types

import (
	"go/ast"
)

type method struct {
	name *ast.Ident
	params []Argument
	resuts []Argument
	structsResolved bool
}
