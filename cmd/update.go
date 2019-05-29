package cmd

import (
	"os"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal/parse"
	"path/filepath"
	"strings"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(updateCommand)
}

// rootCmd represents the base command when called without any subcommands
var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update the project based on the service interface",
	Long: "Godin will read the service.go file and parse the service interface.\n" +
		"Based on that information, all missing endpoints are generated, taking the configuration into account.\n" +
		"After the project was initialized, running this command will generate a basic gRPC server using godin-provided " +
		"default values.\n" +
		"You can use 'godin add' to add different modules to your project.",

	Run: updateCmd,
}

func updateCmd(cmd *cobra.Command, args []string) {

	// project must be initialized
	if _, err := os.Stat(godin.ConfigFilename()); err != nil {
		logrus.Fatal("project not initialized")
	}

	if err := godin.LoadConfiguration(); err != nil {
		logrus.Fatalf("failed to load configuration: %s", err.Error())
	}

	// interface name: Title of ServiceName
	interfaceName := strings.Title(viper.GetString("service.name"))

	// parse service.go
	wd, _ := os.Getwd()
	service := parse.NewServiceParser(filepath.Join(wd, "internal", "service", "service.go"))
	if err := service.ParseFile(); err != nil {
		logrus.Fatalf("failed to parse service.go: %s", err.Error())
	}
	logrus.Info("parsed service.go")

	if err := service.FindInterface(interfaceName); err != nil {
		logrus.Fatalf("unable to find service interface: %s", err.Error())
	}
	logrus.Infof("found interface: %s", interfaceName)

	if err := service.ValidateInterface(); err != nil {
		logrus.Fatalf("interface invalid: %s", err.Error())
	}
	logrus.Info("interface is valid")

	// tplContext := template.NewContextFromConfig()

}
