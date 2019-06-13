package godin

import (
	"os"
	"time"

	"path/filepath"

	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/fs"
	"github.com/lukasjarosch/godin/internal/parse"
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// DefaultDirectoryList holds all of Godin's default paths which are created upon project initialization.
var DefaultDirectoryList = []string{
	"internal/grpc",
	"internal/service",
	"internal/service/endpoint",
	"internal/service/middleware",
	"pkg/grpc",
}

type Project struct {
}

func NewProject() *Project {
	return &Project{}
}

// InitializeConfiguration will setup 'godin.toml' with the default data.
// Then the user will be prompted for additional configuration data.
// At last, the config-file is saved to disk
func (p *Project) InitializeConfiguration() {

	// Godin default config
	config.Set("godin.version", internal.Version)
	config.Set("project.created", time.Now().Format(time.RFC1123))

	config.Set("service.middleware.logging", true)
	config.Set("service.middleware.recovery", true)
	config.Set("service.middleware.authorization", false)
	config.Set("service.middleware.caching", false)

	// prompt for required data and save it to config
	config.Set("service.name", prompt.ServiceName())
	config.Set("service.namespace", prompt.ServiceNamespace())
	config.Set("service.module", prompt.ServiceModule())
	config.Set("protobuf.service", prompt.ProtoServiceName())
	config.Set("protobuf.package", prompt.ProtoPackage())

	SaveConfiguration()
}

// SetupDirectories will setup all directories given in DefaultDirectoryList and also the service
// specific folders: cmd/<serviceName>, internal/service/<serviceName>
func (p *Project) SetupDirectories() error {

	// add service-specific folders
	serviceCmdDir := filepath.Join("cmd", config.GetString("service.name"))
	serviceDir := filepath.Join("internal", "service", config.GetString("service.name"))

	dirs := DefaultDirectoryList
	dirs = append(dirs, serviceCmdDir)
	dirs = append(dirs, serviceDir)

	if err := fs.MakeDirs(dirs); err != nil {
		return errors.Wrap(err, "SetupDirectories")
	}

	return nil
}

// ParseServiceFile will locate the 'service.go', parse it and validate the service interface
// The parsed service is then returned.
func ParseServiceFile(interfaceName string) *parse.Service {
	wd, _ := os.Getwd()
	filePath := filepath.Join("internal", "service", "service.go")

	service := parse.NewServiceParser(filepath.Join(wd, filePath))
	if err := service.ParseFile(); err != nil {
		logrus.Fatalf("failed to parse service.go: %s", err.Error())
	}
	logrus.Debugf("parsed service file: %s", filePath)

	if err := service.FindInterface(interfaceName); err != nil {
		logrus.Fatalf("unable to find service interface: %s", err.Error())
	}
	logrus.Debugf("found service interface: %s", interfaceName)

	if err := service.ValidateInterface(); err != nil {
		logrus.Fatalf("service interface is invalid: %s", err.Error())
	}
	logrus.Debugf("service interface is valid")

	return service
}
