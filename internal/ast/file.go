package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/godin/internal/ast/types"
)

type File struct {
	Filename string
	File     io.Reader
}

func NewFile(filename string, source io.Reader) *File {
	return &File{
		Filename: filename,
		File:     source,
	}
}

func (f *File) Process() error {
	file, err := f.parseFile()
	if err != nil {
	    return err
	}

	ctx, err := f.extractContext(file)
	if err != nil {
	    logrus.Fatal(err)
	}

	logrus.Infof("%+v", ctx)

	return nil
}

func (f *File) parseFile() (*ast.File, error) {
	file, err := parser.ParseFile(token.NewFileSet(), f.Filename, f.File, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *File) extractContext(file ast.Node) (*types.AstContext, error) {
	context := &types.AstContext{}
	visitor := &parseVisitor{src: context}

	ast.Walk(visitor, file)

	return context, nil // TODO: context.validate
}
