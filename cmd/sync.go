package cmd

import (
	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal/project"
	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/generate"
)

func init() {
	rootCmd.AddCommand(syncCommand)
}

// rootCmd represents the base command when called without any subcommands
var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync the generated files with the service interface",
	Run:   syncCmd,
}

func syncCmd(cmd *cobra.Command, args []string) {
	if err := project.HasConfig(); err != nil {
		logrus.Fatal("project not initialized")
	}

	// setup the template data
	data := internal.DataFromConfig()

	l := generate.NewLogging(data)
	l.Render()
}
