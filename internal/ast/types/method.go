package types

import (
	"go/ast"
)

type Method struct {
	name *ast.Ident
	params []Argument
	resuts []Argument
	structsResolved bool
}
