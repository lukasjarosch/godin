package cmd

import (
	"os"

	"github.com/lukasjarosch/godin/internal/module"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal/project"
)

var AvailableModules = []string{"endpoint", "repository", "consumer", "producer", "test"}

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
	project.EnsurePath(ConfigFile)

	// ask user what to do
	mod, err := promptModule()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	switch mod {
	case "endpoint":
		m := module.NewEndpoint()
		err := m.Execute()
		if err != nil {
			logrus.Fatal(err)
		}
		break
	case "repository":
		logrus.Fatal("work in progress")
		break
	case "consumer":
		logrus.Fatal("work in progress")
		break
	case "producer":
		logrus.Fatal("work in progress")
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
		return "", err
	}

	return result, nil
}
