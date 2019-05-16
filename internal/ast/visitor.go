package ast

import (
	"go/ast"

	"github.com/lukasjarosch/godin/internal/ast/types"
)

type (
	parseVisitor struct {
		src *types.AstContext
	}

	typeSpecVisitor struct {
		src  *types.AstContext
		node *ast.TypeSpec
		name *ast.Ident
		iface *types.Iface
	}
)

// Visit implementation for parseVisitor. It will run on the root ast node fetched from Parse()
func (v *parseVisitor) Visit(n ast.Node) ast.Visitor {
	switch rn := n.(type) {
	default:
		return v
	case *ast.File:
		v.src.Pkg = rn.Name
		return v
	case *ast.ImportSpec:
		v.src.Imports = append(v.src.Imports, rn)
		return nil

	case *ast.TypeSpec:
		switch rn.Type.(type) {
		default:
			v.src.Types = append(v.src.Types, rn)
		case *ast.InterfaceType:
			// skip
		}
		return &typeSpecVisitor{src: v.src, node: rn}
	}
}

// Visit implementation for TypeSpecs
func (v *typeSpecVisitor) Visit(n ast.Node) ast.Visitor {
	switch rn := n.(type) {
	default:
		return v
	case *ast.Ident:
		if v.name == nil {
			v.name = rn
		}
		return v
	case *ast.InterfaceType:
		// TODO: return interfaceTypeVisitor
		return nil
	case nil:
		if v.iface != nil {
			v.iface.Name = v.name
			sn := *v.name
			v.iface.Stub = &sn
			v.iface.Stub.Name = v.name.String()
			v.src.Interfaces = append(v.src.Interfaces, *v.iface)
		}
		return nil
	}
}
