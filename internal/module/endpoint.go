package module

import (
	"go/printer"
	"os"

	"github.com/dave/dst/decorator"
	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/ast"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"go/token"
	"go/parser"
	"github.com/dave/dst"
)

type Endpoint struct {
	Data   *internal.Data
	Source *ast.ServiceFileContext
}

func NewEndpoint(data *internal.Data, source *ast.ServiceFileContext) *Endpoint {
	return &Endpoint{Data: data, Source: source}
}

func (e *Endpoint) Execute() error {
	var endpoints []string
	for _, meth := range e.Source.Interfaces[0].Methods {
		endpoints = append(endpoints, meth.Name.Name)
	}

	/*
			endpoint := e.prompt(endpoints)


			var method types.Method


		fset := token.NewFileSet()
		d := decorator.NewDecorator(fset)
		_, err := d.DecorateFile(e.Source.File)
		if err != nil {
			logrus.Fatal(err)
		}
	*/

	// pretend we've generated a partial template (add a tmp package for the parser to work)
	code := `package tmp 
	func Greeting(ctx context.Context, name string) (greeting string, err error ) { return "foo", nil }`
	templateAst, err := decorator.Parse(code)
	if err != nil {
		logrus.Fatal(err)
	}

	// dst => ast
	_, newFile, err := decorator.RestoreFile(templateAst)
	if err != nil {
		logrus.Fatal(err)
	}

	// write
	targetFile, _ := os.Create("/tmp/test.go")
	defer targetFile.Close()
	if err := printer.Fprint(targetFile, e.Source.FSet, newFile); err != nil {
		logrus.Fatal(err)
	}

	// reopen  and parse (pretend it's the output file)
	tmpf, _ := os.Open("/tmp/test.go")
	tmpdec, err := decorator.ParseFile(token.NewFileSet(), "foo.go", tmpf, parser.DeclarationErrors)
	if err != nil {
	    logrus.Fatal(err)
	}

	dst.Fprint(os.Stdout, tmpdec.Decls[0].(*dst.FuncDecl), nil)

	if err := printer.Fprint(targetFile, e.Source.FSet, newFile); err != nil {
		logrus.Fatal(err)
	}

	return nil
}

func (e *Endpoint) prompt(selection []string) string {
	logrus.Info("Before adding a new endpoint, extend the service interface")
	prompt := promptui.Select{
		Label: "Select endpoint to add",
		Items: selection,
	}

	_, result, err := prompt.Run()
	if err != nil {
		logrus.Info("Bye...")
		os.Exit(1)
	}
	return result
}
