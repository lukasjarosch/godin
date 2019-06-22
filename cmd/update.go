package cmd

import (
	"fmt"
	"os"

	"strings"

	"path"

	"github.com/lukasjarosch/godin/internal/generate"
	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	config "github.com/spf13/viper"
	"github.com/lukasjarosch/godin/internal/fs"
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

	logrus.SetLevel(logrus.DebugLevel)

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

	// update internal/service/<serviceName>/implementation.go
	implementationGen := generate.NewImplementation(TemplateFilesystem, service.Interface, tplContext)
	if err := implementationGen.Update(); err != nil {
		logrus.Errorf("failed to update implementation: %s: %s", implementationGen.TargetPath(), err.Error())
	} else {
		logrus.Infof("updated implementation: %s", implementationGen.TargetPath())
	}

	// request_response.go
	reqRes := generate.NewRequestResponse(TemplateFilesystem, service.Interface, tplContext)
	if err := reqRes.Update(); err != nil {
		logrus.Error(fmt.Sprintf("failed to generate request_response.go: %s", err.Error()))
	} else {
		logrus.Infof("generated %s", reqRes.TargetPath())
	}

	// set.go
	endpointSet := generate.NewEndpointSet(TemplateFilesystem, service.Interface, tplContext)
	if err := endpointSet.Update(); err != nil {
		logrus.Errorf("failed to update endpoint set: %s: %s", endpointSet.TargetPath(), err)
	} else {
		logrus.Infof("updated endpoint set: %s", endpointSet.TargetPath())
	}

	// endpoints.go
	endpoints := generate.NewEndpoints(TemplateFilesystem, service.Interface, tplContext)
	if err := endpoints.Update(); err != nil {
		logrus.Errorf("failed to update endpoints %s: %s", endpoints.TargetPath(), err)
	} else {
		logrus.Infof("updated endpoints: %s", endpoints.TargetPath())
	}

	// middleware.go
	middleware := generate.NewMiddleware(TemplateFilesystem, service.Interface, tplContext)
	if err := middleware.Update(); err != nil {
		logrus.Error(fmt.Sprintf("failed to generate middleware.go: %s", err.Error()))
	} else {
		logrus.Infof("generated %s", middleware.TargetPath())
	}

	// logging middleware
	if config.GetBool("service.middleware.logging") {
		logging := generate.NewLoggingMiddleware(TemplateFilesystem, service.Interface, tplContext)
		if err := logging.Update(); err != nil {
			logrus.Errorf("failed to update logging middleware: %s: %s", logging.TargetPath(), err.Error())
		} else {
			logrus.Infof("updated logging middleware: %s", logging.TargetPath())
		}
	}

	if config.GetBool("grpc.enabled") {
		// grpc/request_response.go
		grpcRequestResponse := generate.NewGrpcRequestResponse(TemplateFilesystem, service.Interface, tplContext)
		fs.MakeDirs([]string{path.Dir(grpcRequestResponse.TargetPath())}) // ignore errors, just ensure the path exists
		if err := grpcRequestResponse.Update(); err != nil {
			logrus.Errorf("failed to update grpc/request_response.go: %s", err)
		} else {
			logrus.Infof("updated %s", grpcRequestResponse.TargetPath())
		}

		// grpc/encode_decode.go
		grpcEncodeDecode := generate.NewGrpcEncodeDecode(TemplateFilesystem, service.Interface, tplContext)
		if err := grpcEncodeDecode.Update(); err != nil {
			logrus.Errorf("failed to update grpc/encode_decode.go: %s", err)
		} else {
			logrus.Infof("updated %s", grpcEncodeDecode.TargetPath())
		}
	}
}
