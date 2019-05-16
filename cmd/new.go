package cmd

import (
	"fmt"
	"os"

	"path"

	"github.com/lukasjarosch/godin/internal/project"
	prompting "github.com/lukasjarosch/godin/internal/prompt"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

// rootCmd represents the base command when called without any subcommands
var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Setup a new microservice project",
	Run:   handler,
}

func handler(cmd *cobra.Command, args []string) {

	logrus.SetLevel(logrus.DebugLevel)
	project.EnsureConfig(false)

	viper.Set("godin.version", Version)
	viper.Set("godin.commit", Commit)
	viper.Set("godin.build", BuildDate)

	// initialize config with user data
	projectPath, _ := os.Getwd()
	viper.Set("project.path", projectPath)
	prompt()
	project.SaveConfig()

	// setup the template data
	data := &template.Data{
		Project: template.Project{
			RootPath: viper.GetString("project.path"),
		},
		Godin: template.Godin{
			Version:Version,
			Commit:Commit,
			Build:BuildDate,
		},
		Service: template.Service{
			Name: viper.GetString("service.name"),
			Namespace: viper.GetString("service.namespace"),
			Module: viper.GetString("service.module"),
		},
	}

	// set-up godin project
	godin := project.NewGodinProject(data, Box)

	// add all required folders
	godin.AddFolder("cmd")
	godin.AddFolder("internal")
	godin.AddFolder("pkg")
	godin.AddFolder("deployment")
	godin.AddFolder("deployment/k8s")
	godin.AddFolder("internal/grpc")
	godin.AddFolder("internal/service")
	godin.AddFolder(fmt.Sprintf("internal/service/%s", viper.GetString("service.name")))
	godin.AddFolder("internal/service/middleware")
	godin.AddFolder("internal/service/endpoint")
	if err := godin.MkdirAll(); err != nil {
		logrus.Fatal(err)
	}

	// add some basic templates
	godin.AddTemplate(template.NewTemplateFile("README.tpl", path.Join(projectPath, "README.md"), false))
	godin.AddTemplate(template.NewTemplateFile("gitignore.tpl", path.Join(projectPath, ".gitignore"), false))
	godin.AddTemplate(template.NewTemplateFile("Dockerfile.tpl", path.Join(projectPath, "Dockerfile"), false))

	godin.AddTemplate(template.NewTemplateFile("./internal/service.tpl", path.Join(projectPath, "internal", "service.go"), true))
	godin.AddTemplate(template.NewTemplateFile("./internal/models.tpl", path.Join(projectPath, "internal", "models.go"), true))
	godin.AddTemplate(template.NewTemplateFile(
		"./internal/service/service/implementation.tpl",
		path.Join(projectPath, "internal", "service", viper.GetString("service.name"), "implementation.go"),
		true))
	godin.AddTemplate(template.NewTemplateFile(
		"./internal/service/middleware/middleware.tpl",
		path.Join(projectPath, "internal", "service", "middleware", "middleware.go"),
		true))

	godin.Render()

	// initialize module
	godin.InitModule(viper.GetString("service.module"))
}

// prompt the user for all required values and store them in viper
func prompt() {

	// service.name
	prompt := prompting.NewPrompt(
		"Enter the service name (lowercase)",
		"",
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.Lowercase(),
		),
	)
	serviceName, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("service.name", serviceName)

	// service.namespace
	prompt = prompting.NewPrompt(
		"Enter the namespace (lowercase)",
		"",
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.Lowercase(),
		),
	)
	namespace, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("service.namespace", namespace)

	// service.module
	prompt = prompting.NewPrompt(
		"Enter the go-module",
		fmt.Sprintf("bitbucket.org/jdbergmann/%s/%s", namespace, serviceName),
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.Lowercase(),
		),
	)
	module, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("service.module", module)

	// protobuf.service
	prompt = prompting.NewPrompt(
		"Enter the gRPC service name which you want to implement (CamelCase)",
		"",
		prompting.Validate(
			prompting.MinLengthThree(),
		),
	)
	grpcService, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("protobuf.service", grpcService)

	// protobuf.package
	prompt = prompting.NewPrompt(
		"Enter the protobuf package of the gRPC service",
		"",
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.Lowercase(),
		),
	)
	protoPackage, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("protobuf.package", protoPackage)
}
