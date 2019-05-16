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
		src   *types.AstContext
		node  *ast.TypeSpec
		name  *ast.Ident
		iface *types.Iface
	}

	interfaceTypeVisitor struct {
		node    *ast.TypeSpec
		ts      *typeSpecVisitor
		methods []types.Method
	}

	methodVisitor struct {
		depth           int
		node            *ast.TypeSpec
		list            *[]types.Method
		name            *ast.Ident
		params, results *[]types.Argument
		isMethod        bool
	}

	argListVisitor struct {
		list *[]types.Argument
	}

	argVisitor struct {
		node  *ast.TypeSpec
		parts []ast.Expr
		list  *[]types.Argument
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
		return &interfaceTypeVisitor{ts: v, methods: []types.Method{}}
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

func (v *interfaceTypeVisitor) Visit(n ast.Node) ast.Visitor {
	switch n.(type) {
	default:
		return v
	case *ast.Field:
		return &methodVisitor{list: &v.methods}
	case nil:
		v.ts.iface = &types.Iface{Methods: v.methods}
		return nil
	}
}

func (v *methodVisitor) Visit(n ast.Node) ast.Visitor {
	switch rn := n.(type) {
	default:
		v.depth++
		return v
	case *ast.Ident:
		if rn.IsExported() {
			v.name = rn
		}
		return v
	case *ast.FieldList:
		if v.params == nil {
			v.params = &[]types.Argument{}
			return &argListVisitor{list: v.params}
		}
		if v.results == nil {
			v.results = &[]types.Argument{}
		}
		return &argListVisitor{list: v.results}
	}
}

func (v *argListVisitor) Visit(n ast.Node) ast.Visitor {
	switch n.(type) {
	default:
		return nil
	case *ast.Field:
		return &argVisitor{list: v.list}
	}
}

func (v *argVisitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.CommentGroup, *ast.BasicLit:
		return nil
	case *ast.Ident:
		if t.Name != "_" {
			v.parts = append(v.parts, t)
		}
	case ast.Expr:
		v.parts = append(v.parts, t)
	case nil:
		names := v.parts[:len(v.parts)-1]
		tp := v.parts[len(v.parts)-1]
		if len(names) == 0 {
			*v.list = append(*v.list, types.Argument{Typ: tp})
			return nil
		}
		for _, n := range names {
			*v.list = append(*v.list, types.Argument{
				Name: n.(*ast.Ident),
				Typ: tp,
			})
		}
	}
	return nil
}