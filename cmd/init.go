package cmd

import (
	"os"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal/template"
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

	if err := template.WriteDockerfile(TemplateFilesystem); err != nil {
		logrus.Error(err)
	}
	logrus.Info("generated Dockerfile")

	if err := template.WriteGitignore(TemplateFilesystem); err != nil {
		logrus.Error(err)
	}
	logrus.Info("generated .gitignore")
}
