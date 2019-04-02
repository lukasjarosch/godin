package cmd

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/project"
	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

// rootCmd represents the base command when called without any subcommands
var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Setup a new microservice structure",
	Run:   handler,
}

func handler(cmd *cobra.Command, args []string) {

	logrus.SetLevel(logrus.DebugLevel)

	// load spec
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

	err = spec.ResolveDependencies()
	if err != nil {
	    logrus.Fatal(err)
	}

	// setup new project with specification
	godin := project.NewGodinProject(spec, projectPath)

	// add all required folders
	godin.AddFolder("internal")
	godin.AddFolder("cmd")
	godin.AddFolder(path.Join("cmd", spec.Service.Name))
	godin.AddFolder("k8s")
	godin.AddFolder("internal/server")
	godin.AddFolder(path.Join("internal", spec.Service.Name))
	godin.AddFolder("internal/config")

	if err := godin.MkdirAll(); err != nil {
		logrus.Fatal(err)
	}

	// add some basic templates
	godin.AddTemplate(template.NewTemplateFile("README.tpl", path.Join(projectPath, "README.md"), false))
	godin.AddTemplate(template.NewTemplateFile("gitignore.tpl", path.Join(projectPath, ".gitignore"), false))
	godin.AddTemplate(template.NewTemplateFile("Dockerfile.tpl", path.Join(projectPath, "Dockerfile"), false))

	godin.AddTemplate(template.NewTemplateFile("config.tpl", path.Join(projectPath, "internal", "config", "config.go"), true))
	godin.AddTemplate(template.NewTemplateFile("service.tpl", path.Join(projectPath, "internal", spec.Service.Name, "service.go"), true))
	godin.AddTemplate(template.NewTemplateFile("models.tpl", path.Join(projectPath, "internal", spec.Service.Name, "models.go"), true))
	godin.AddTemplate(template.NewTemplateFile("handler.tpl", path.Join(projectPath, "internal", "server", "handler.go"), true))
	godin.AddTemplate(template.NewTemplateFile("server.tpl", path.Join(projectPath, "internal", "server", "server.go"), true))
	godin.AddTemplate(template.NewTemplateFile("main.tpl", path.Join(projectPath, "cmd", spec.Service.Name, "main.go"), true))

	godin.Render()
}
