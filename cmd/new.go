package cmd

import (
	"os"
	"path"

	"github.com/lukasjarosch/godin/internal/project"
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

	serviceName := "greeter"

	projectPath, _ := os.Getwd()
	projectPath = path.Join(projectPath, "examples", serviceName)

	// create a bare-bones Godin project
	godin := project.NewGodinProject(serviceName, projectPath)

	// add all required folders
	godin.AddFolder("internal")
	godin.AddFolder("cmd")
	godin.AddFolder(path.Join("cmd", serviceName))
	godin.AddFolder("k8s")
	godin.AddFolder("internal/server")
	godin.AddFolder("internal/service")
	godin.AddFolder("internal/config")

	if err := godin.MkdirAll(); err != nil {
		logrus.Fatal(err)
	}

	// add some basic templates
	godin.AddTemplate(template.NewTemplateFile("README.tpl", path.Join(projectPath, "README.md"), false))
	godin.AddTemplate(template.NewTemplateFile("gitignore.tpl", path.Join(projectPath, ".gitignore"), false))
	godin.AddTemplate(template.NewTemplateFile("Dockerfile.tpl", path.Join(projectPath, "Dockerfile"), false))

	godin.Render()

}
