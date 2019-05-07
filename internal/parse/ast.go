package parse

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

type astParser struct {
}

func FromString(input string) (*dst.File, error) {

	input = "package tmp\n" + input

	f, err := decorator.Parse(input)
	if err != nil {
	    return nil, err
	}

	return f, nil

	/*
	if !strings.Contains(input, "package") {
		input = "package main\n" + input
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tmp.go", input, parser.ParseComments)
	if err != nil {
	    return nil, err
	}

	return f, nil
	*/
}
