package types

import (
	"go/ast"
)

type Iface struct {
	Name, Stub, Receiver *ast.Ident
	Methods []Method
}