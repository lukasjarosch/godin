package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"

	"github.com/sirupsen/logrus"
)

type File struct {
	Filename string
	File     io.Reader
	FSet *token.FileSet
}

func NewFile(filename string, source io.Reader) *File {
	return &File{
		Filename: filename,
		File:     source,
	}
}

func (f *File) Process() (ctx *ServiceFileContext, err error) {

	// parse AST of file
	file, err := f.parseFile()
	if err != nil {
		return nil, err
	}

	// parse the service interface and assemble the ServiceFileContext
	serviceFileContext, err := f.extractContext(file)
	if err != nil {
		logrus.Fatal(err)
	}

	serviceFileContext.File = file
	serviceFileContext.FSet = f.FSet

	return serviceFileContext, nil
}

func (f *File) parseFile() (*ast.File, error) {
	f.FSet = token.NewFileSet()
	file, err := parser.ParseFile(f.FSet, f.Filename, f.File, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *File) extractContext(file ast.Node) (*ServiceFileContext, error) {
	context := &ServiceFileContext{}
	visitor := &parseVisitor{src: context}

	ast.Walk(visitor, file)

	return context, context.Validate()
}
