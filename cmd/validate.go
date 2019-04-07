package cmd

import (
	"go/parser"
	"go/token"
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/lukasjarosch/godin/internal/validate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vetcher/go-astra"
)

func init() {
	rootCmd.AddCommand(validateCommand)
}

var validateCommand = &cobra.Command{
	Use:   "validate",
	Short: "Validate a service against it's specification",
	Run:   validateHandler,
}

func validateHandler(cmd *cobra.Command, args []string) {
	spec := loadSpec()
	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, "examples/spec-greeter/internal/greeter/service.go", nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		logrus.Fatal(err)
	}

	file, err := astra.ParseAstFile(node)
	if err != nil {
		logrus.Fatal(err)
	}

	p := validate.NewParser(file)

	// find missing methods
	for _, meth := range spec.Service.Methods {
		if !p.HasMethod(meth.Name) {
			logrus.Errorf("MISSING: %s", meth.Signature())
		}
	}

	// ensure the handler specification matches the service implementation
	for _, s := range file.Structures {
		if spec.Service.Handler == s.Name {
			for _, meth := range s.Methods {
				if !spec.Service.HasMethod(meth.Name) {
					logrus.Warnf("UNSPECIFIED: %s", meth.Function.String())
					continue
				}

				if err := validate.ValidateMethod(s.Name, spec.Service.GetMethod(meth.Name), meth); err != nil {
					logrus.Error(err.Error())
				} else {
					logrus.Infof("VALID: %s", meth.Function.String())
				}
			}
		}
	}
}

func loadSpec() *specification.Specification {
	cwd, _ := os.Getwd()
	projectPath := path.Join(cwd, "examples", "spec-greeter")
	spec, err := specification.LoadPath(path.Join(projectPath, "greeter.yaml"))
	if err != nil {
		logrus.Fatalf("failed to load specification: %v", err)
	}

	err = spec.Validate()
	if err != nil {
		logrus.Fatal(err)
	}

	return spec
}
