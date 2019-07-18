package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/lukasjarosch/godin/internal/bundle"
	"github.com/lukasjarosch/godin/internal/bundle/transport"
	"github.com/lukasjarosch/godin/internal/godin"
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

// rootCmd represents the base command when called without any subcommands
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add bundles to an initalized godin project",
	Long: `Add different kind of bundles to the current project. 

	godin add <bundleType>

The available bundles are:  
  middleware: to add a new service-layer middleware 
  datastore: to add a new datastore bundle (mysql, mongodb, redis)
  subscriber: add an AMQP topic subscription
  transport: to add a new transport layer`,
	Run:  addCmd,
	Args: cobra.MinimumNArgs(1),
}

var validBundleTypes = []string{"middleware", "datastore", "transport", "subscriber", "publisher"}

func addCmd(cmd *cobra.Command, args []string) {
	logrus.SetLevel(logrus.DebugLevel)
	// ensure a valid bundle type is passed
	bundleType, err := validateBundleType(args[0])
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	// project must be initialized
	if _, err := os.Stat(godin.ConfigFilename()); err != nil {
		logrus.Fatal("project not initialized")
	}

	if err := godin.LoadConfiguration(); err != nil {
		logrus.Fatalf("failed to load configuration: %s", err.Error())
	}

	switch bundleType {
	case "middleware",
		"datastore":
		logrus.Info("sorry, this bundle type is not yet implemented :(")
	case "transport":
		selected, err := prompt.NewSelect("Which transport layer do you want to add?", []string{"amqp", "grpc", "http"})
		if err != nil {
			logrus.Errorf("select cancelled: %s", err)
			os.Exit(1)
		}
		switch selected {
		case "amqp":
			_, err := transport.InitializeAMQP()
			if err != nil {
				logrus.Errorf("failed to initialize AMQP transport: %s", err)
			}
			// TODO: generate

		case "grpc",
			"http":
			logrus.Info("not yet implemented")
			os.Exit(1)
		}

	case "subscriber":
		_, err := transport.InitializeAMQP()
		if err != nil {
			logrus.Errorf("failed to initialize AMQP transport: %s", err)
		}

		_, err = bundle.InitializeSubscriber()
		if err != nil {
			logrus.Errorf("failed to initialize subscriber: %s", err)
			os.Exit(1)
		}

	case "publisher":
		_, err := transport.InitializeAMQP()
		if err != nil {
			logrus.Errorf("failed to initialize AMQP transport: %s", err)
		}

		_, err = bundle.InitializePublisher()
		if err != nil {
			logrus.Errorf("failed to initialize publisher: %s", err)
			os.Exit(1)
		}


		// TODO: godin.json is NOT a service configuration, thus the 'topic', 'queue' and 'exchange' values must be configurable with ENV variables
	}

	logrus.Info("The next time you run 'godin update', the new bundles are generated into your service.")
}

func validateBundleType(givenType string) (bundleType string, err error) {
	for _, validType := range validBundleTypes {
		if givenType == validType {
			bundleType = validType
			break
		}
	}
	if bundleType == "" {
		return "", fmt.Errorf("invalid bundle type, valid types are: %s", strings.Join(validBundleTypes, ", "))
	}
	return bundleType, nil
}
