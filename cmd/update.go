package cmd

import (
	"os"

	"strings"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	config "github.com/spf13/viper"
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

	// parse service.go
	interfaceName := strings.Title(config.GetString("service.name"))
	service := godin.ParseServiceFile(interfaceName)

	// prepare template context for rendering
	tplContext := template.NewContextFromConfig()
	tplContext = template.PopulateFromService(tplContext, service)

	logrus.Infof("%+v", tplContext)

}
