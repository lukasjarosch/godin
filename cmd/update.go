package cmd

import (
	"os"

	"strings"

	"path/filepath"

	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/parse"
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

	implementationFile := filepath.Join("internal", "service", tplContext.Service.Name, "implementation.go")

	if _, err := os.Stat(implementationFile); err != nil {
		GenerateFullImplementation(tplContext, implementationFile)
	} else {
		logrus.Info("implementation.go already exist, updating")

		implementation := parse.NewImplementationParser(implementationFile, service.Interface)
		if err := implementation.Parse(); err != nil {
			logrus.Fatalf("unable to parse implementation.go: %s", err.Error())
		}
		logrus.Infof("parsed %s", implementationFile)

		if len(implementation.MissingMethods) > 0 {
			for _, meth := range implementation.MissingMethods {
				logrus.Infof("missing method: %s", meth.String())

				tpl := template.NewPartial("service_method", true)
				data, err := tpl.Render(TemplateFilesystem, template.MethodFromType(meth))
				if err != nil {
					logrus.Fatalf("failed to render partial template: %s", err.Error())
				}

				writer := template.NewFileAppendWriter(implementationFile, data)
				if err := writer.Write(); err != nil {
					logrus.Fatalf("failed to write file: %s", implementationFile)
				}
				logrus.Info("updated implementation")
			}
		} else if len(implementation.File.Methods) > len(service.Interface.Methods) {
			logrus.Info("there are too many methods, i cannot remove them, you need to to that!")
			logrus.Info("only count exported methods and find which one needs to be removed by the developer")
		} else {
			logrus.Info("all methods of the interface are present")
		}
	}
}

func GenerateFullImplementation(tplContext template.Context, targetPath string) {
	logrus.Info("implementation.go does not yet exist, creating")

	implementation := template.NewGenerator(template.ImplementationFileOptions(tplContext, targetPath))
	if err := implementation.GenerateFile(TemplateFilesystem); err != nil {
		logrus.Fatalf("failed to generate implementation.go: %s", err.Error())
	}

	logrus.Infof("generated %s", targetPath)
}
