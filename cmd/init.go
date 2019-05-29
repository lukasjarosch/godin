package cmd

import (
	"os"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal/template"
	"fmt"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

// rootCmd represents the base command when called without any subcommands
var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new godin microservice project in the current directory",
	Run:   initCmd,
}

func initCmd(cmd *cobra.Command, args []string) {

	logrus.SetLevel(logrus.DebugLevel)

	// check if already an initialized project
	if _, err := os.Stat(godin.ConfigFilename()); err == nil {
		logrus.Fatal("project already initialized")
	}

	project := godin.NewProject()
	project.InitializeConfiguration()

	if err := project.SetupDirectories(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("folder structure created")

	// service.go
	service := template.NewGenerator(template.ServiceStubOptions())
	if err := service.GenerateFile(TemplateFilesystem); err != nil {
		logrus.Error(fmt.Sprintf("failed to generate service stub: %s", err.Error()))
	} else {
		logrus.Info("generated internal/service/service.go")
	}

	// Dockerfile
	dockerfile := template.NewGenerator(template.DockerfileOptions())
	if err := dockerfile.GenerateFile(TemplateFilesystem); err != nil {
		logrus.Error(fmt.Sprintf("failed to generate Dockerfile: %s", err.Error()))
	} else {
		logrus.Info("generated Dockerfile")
	}

	// gitignore
	gitignore := template.NewGenerator(template.GitignoreOptions())
	if err := gitignore.GenerateFile(TemplateFilesystem); err != nil {
		logrus.Error(fmt.Sprintf("failed to generate .gitinore: %s", err.Error()))
	} else {
		logrus.Info("generated .gitinore")
	}
}
