package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"strings"

	"github.com/lukasjarosch/godin/internal/project"
	prompting "github.com/lukasjarosch/godin/internal/prompt"
	"github.com/lukasjarosch/godin/internal/specification"
	"github.com/lukasjarosch/godin/internal/template"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

const ConfigFile = "godin.toml"

// rootCmd represents the base command when called without any subcommands
var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Setup a new microservice structure",
	Run:   handler,
}

func handler(cmd *cobra.Command, args []string) {

	logrus.SetLevel(logrus.DebugLevel)

	// create config file and load it directly
	if _, err := os.Stat(ConfigFile); err != nil {
		os.Create(ConfigFile)
	}
	viper.SetConfigFile(ConfigFile)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	// prompting
	prompt := prompting.NewPrompt(
		"Enter the service name (lowercase):",
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

	prompt = prompting.NewPrompt(
		"Enter the namespace (lowercase):",
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

	prompt = prompting.NewPrompt(
		"Enter the service-file filename",
		fmt.Sprintf("%s.go", serviceName),
		prompting.Validate(
			prompting.MinLengthThree(),
			prompting.Lowercase(),
			prompting.GoSuffix(),
		),
	)
	serviceFile, err := prompt.Run()
	if err != nil {
		os.Exit(1)
	}
	viper.Set("service.file", serviceFile)

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

	logrus.Info(serviceName)
	logrus.Info(namespace)
	logrus.Info(serviceFile)
	logrus.Info(module)

	viper.WriteConfigAs(ConfigFile)

	return

	// load spec
	cwd, _ := os.Getwd()
	projectPath := path.Join(cwd, "examples", "spec-greeter")
	spec, err := specification.LoadPath(path.Join(projectPath, "greeter.yaml"))
	if err != nil {
		logrus.Fatalf("failed to load specification: %v", err)
	}

	err = spec.Validate()
	if err != nil {
		logrus.Fatal(err)
	}

	err = spec.ResolveDependencies()
	if err != nil {
		logrus.Fatal(err)
	}

	// setup new project with specification
	godin := project.NewGodinProject(spec, projectPath)

	// add all required folders
	godin.AddFolder("internal")
	godin.AddFolder("cmd")
	godin.AddFolder(path.Join("cmd", spec.Service.Name))
	godin.AddFolder("k8s")
	godin.AddFolder("internal/server")
	godin.AddFolder(path.Join("internal", spec.Service.Name))
	godin.AddFolder("internal/config")

	if err := godin.MkdirAll(); err != nil {
		logrus.Fatal(err)
	}

	// add some basic templates
	godin.AddTemplate(template.NewTemplateFile("README.tpl", path.Join(projectPath, "README.md"), false))
	godin.AddTemplate(template.NewTemplateFile("gitignore.tpl", path.Join(projectPath, ".gitignore"), false))
	godin.AddTemplate(template.NewTemplateFile("Dockerfile.tpl", path.Join(projectPath, "Dockerfile"), false))

	godin.AddTemplate(template.NewTemplateFile("config.tpl", path.Join(projectPath, "internal", "config", "config.go"), true))
	godin.AddTemplate(template.NewTemplateFile("service.tpl", path.Join(projectPath, "internal", spec.Service.Name, "service.go"), true))
	godin.AddTemplate(template.NewTemplateFile("models.tpl", path.Join(projectPath, "internal", spec.Service.Name, "models.go"), true))
	godin.AddTemplate(template.NewTemplateFile("handler.tpl", path.Join(projectPath, "internal", "server", "handler.go"), true))
	godin.AddTemplate(template.NewTemplateFile("server.tpl", path.Join(projectPath, "internal", "server", "server.go"), true))
	godin.AddTemplate(template.NewTemplateFile("main.tpl", path.Join(projectPath, "cmd", spec.Service.Name, "main.go"), true))

	godin.Render()
}

// prompt for value, only lowercase values
func promptLowercase(label, defaultValue string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("required field must be set")
		}
		if strings.ToLower(input) != input {
			return errors.New("lowercase only")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
		Default:  defaultValue,
	}

	data, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

// prompt for a value
func prompt(label string, defaultValue string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("required field must be set")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
		Default:  defaultValue,
	}

	data, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
