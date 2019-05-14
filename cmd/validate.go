package cmd

import (
	"go/parser"
	"go/token"
	"os"
	"path"

	"strings"
	"github.com/dave/dst/decorator"

	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/lukasjarosch/godin/internal/validate"
	"github.com/manifoldco/promptui"
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
	serviceFile := "examples/spec-greeter/internal/greeter/service.go"
	spec := loadSpec()
	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, serviceFile, nil, parser.ParseComments|parser.AllErrors)
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

			//
			prompt := promptui.Prompt{
				Label:     "Should I add this method to your service?",
				IsConfirm: true,
			}

			result, err := prompt.Run()
			if err != nil {
				return
			}

			if strings.ToLower(result) == "y" {
				logrus.Info(meth.Signature())

				m := template.NewMethodTemplate(spec, meth)
				out, err := m.Render()
				if err != nil {
					logrus.Error(err)
					continue
				}
				//logrus.Infof("Okay, here is what I'm going to insert into %s\n%s", serviceFile, out)

				dst, err := parse.FromString(out)
				if err != nil {
					logrus.Error(err)
					continue
				}

				decorator.Print(dst)
			}
		}
	}

	// validate errors

	// ensure the handler specification matches the service implementation
	for _, s := range file.Structures {
		if spec.Service.Handler == s.Name {

			// validate methods
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
