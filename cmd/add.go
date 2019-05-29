package cmd

import (
	"os"

	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/module"
	"github.com/lukasjarosch/godin/internal/project"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AvailableModules = []string{"endpoint", "middleware", "repository", "consumer", "producer", "test"}

func init() {
	rootCmd.AddCommand(addCommand)
}

// rootCmd represents the base command when called without any subcommands
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add endpoints and modules",
	Run:   addCmd,
}

func addCmd(cmd *cobra.Command, args []string) {
	if err := project.HasConfig(); err != nil {
		logrus.Fatal("project not initialized")
	}

	// setup the template data
	data := internal.DataFromConfig()

	// parse service file
	f, err := os.Open(data.ServiceFilePath())
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(f.Name())

	// ask user what to do
	mod, err := promptModule()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	switch mod {
	case "endpoint":
		logrus.Fatal("work in progress")
		break
	case "middleware":
		logrus.Fatal("work in progress")
		break
	case "repository":
		logrus.Fatal("work in progress")
		break
	case "consumer":
		logrus.Fatal("work in progress")
		break
	case "producer":
		m := module.NewProducer()
		if err := m.Execute(); err != nil {
			logrus.Fatal(err)
		}
		break
	case "test":
		logrus.Fatal("work in progress")
		break
	}

}

// promptModule will present a list of all available modules and ask the user to select one.
// The selected module is returned as string.
func promptModule() (string, error) {
	prompt := promptui.Select{
		Label: "What do you want to add?",
		Items: AvailableModules,
	}

	_, result, err := prompt.Run()
	if err != nil {
		logrus.Info("Bye...")
		os.Exit(1)
	}

	return result, nil
}
