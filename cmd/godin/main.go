package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lukasjarosch/godin/internal/godin"
	prompting "github.com/lukasjarosch/godin/internal/prompt"
	"github.com/lukasjarosch/godin/internal/template"

	"github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

var tplContext template.Context

func init() {
	tplContext = template.Context{
		Godin: template.Godin{
			Version: "2.0.0",
		},
		Service: template.Service{
			Name:      "Greeter",
			Namespace: "Godin",
			Module:    "github.com/lukasjarosch/godin/example/greeter/internal/service",
			Methods: []template.Method{
				{
					ServiceName: "Greeter",
					Name:        "DoSomething",
					Params: []template.Variable{
						{
							Name: "ctx",
							Type: "tplContext.Context",
						},
						{
							Name: "thaName",
							Type: "string",
						},
						{
							Name: "language",
							Type: "string",
						},
					},
					Returns: []template.Variable{
						{
							Name: "greeting",
							Type: "string",
						},
						{
							Name: "err",
							Type: "error",
						},
					},
				},
			},
		},
	}
}

func usage() {
	fmt.Println("usage: godin <command> [<args>]")
	fmt.Println()
	fmt.Println("start working on a new microservice")
	fmt.Println("    init           Initialize a new godin microservice project in the current directory")
	fmt.Println("    generate       Generate boilerplate from the main service interface (use after init)")
	fmt.Println()
	fmt.Println("work on an existing project")
	fmt.Println("    update         Parse service interface and generate everything that's missing")
	fmt.Println("    add            Add new modules to the current project (middleware, transport, database integration, ...)")
	fmt.Println("    readme         Generate README from the current implementation")
	fmt.Println("    k8s            Generate Kubernetes service and deployment manifests")
	fmt.Println()
	fmt.Println("miscellaneous functionality")
	fmt.Println("    git-hook       Install git-hooks for this project")
	fmt.Println("    git-unhook     Remove installed git-hooks")
	fmt.Println("    version        Print the Godin version")
	fmt.Println("    update         Update the Godin binary")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "init":
		// check if already an initialized project
		if _, err := os.Stat(godin.ConfigFilename()); err == nil {
			logrus.Fatal("project already initialized")
		}

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

		// service.module
		prompt = prompting.NewPrompt(
			"Enter the Go module name",
			"",
			prompting.Validate(
				prompting.MinLengthThree(),
			),
		)
		module, err := prompt.Run()
		if err != nil {
			os.Exit(1)
		}

		// protobuf.service
		prompt = prompting.NewPrompt(
			"Enter the gRPC service name which this service implements",
			"",
			prompting.Validate(
				prompting.MinLengthThree(),
			),
		)
		protoService, err := prompt.Run()
		if err != nil {
			os.Exit(1)
		}

		// protobuf.package
		prompt = prompting.NewPrompt(
			"Enter the protobuf package of the service",
			"",
			prompting.Validate(
				prompting.MinLengthThree(),
			),
		)
		protoPackage, err := prompt.Run()
		if err != nil {
			os.Exit(1)
		}

		// ensure config file exists and load
		os.Create(godin.ConfigFilename())
		logrus.Debug("config file created")
		if err := godin.LoadConfiguration(); err != nil {
			logrus.Fatalf("failed to load config: %s", err.Error())
		}

		// initialize configuration
		config.Set("project.created", time.Now().Format(time.RFC1123))
		config.Set("service.name", serviceName)
		config.Set("service.namespace", namespace)
		config.Set("service.module", module)
		config.Set("service.middleware.logging", true)
		config.Set("service.middleware.recovery", true)
		config.Set("service.middleware.authorization", true)
		config.Set("service.middleware.caching", true)
		config.Set("protobuf.service", protoService)
		config.Set("protobuf.package", protoPackage)
		if err := godin.SaveConfiguration(); err != nil {
			logrus.Fatal(err)
		}

		// TODO: set after the service file is created
		//config.Set("service.file", "internal/service/service.go")

		logrus.Info("project initialized")
		break

	case "version":
		fmt.Printf("godin version %s\n", "2.0.0")
	default:
		logrus.Fatalf("%q is not a valid command", os.Args[1])
	}

	return

	// render partial
	tpl := template.NewPartial("endpoint", true)
	data, err := tpl.Render(tplContext.Service.Methods[0])
	if err != nil {
		logrus.Fatal(err)
	}

	// apend partial to existing file
	writer := template.NewFileAppendWriter("/tmp/test.go", data)
	if err := writer.Write(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("appended new enpoint to file")

	// README
	tpl2 := template.NewFile("readme", false)
	data, err = tpl2.Render(tplContext)
	if err != nil {
		logrus.Fatal(err)
	}
	writer2 := template.NewFileWriter("/tmp/README.md", data)
	if err := writer2.Write(true); err != nil {
		logrus.Fatal(err)
	}

}

func fullTemplate() {

	// render a full template which imports partials
	tpl := template.NewFile("endpoints", true)
	data, err := tpl.Render(tplContext)
	if err != nil {
		logrus.Fatal(err)
	}

	// write the generated code into a file (overwrite existing file)
	writer := template.NewFileWriter("/tmp/test.go", data)
	if err := writer.Write(true); err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("written file %s", writer.Path)
	}
}
