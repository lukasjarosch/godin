package godin

import (
	"time"

	"path/filepath"

	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/fs"
	"github.com/lukasjarosch/godin/internal/prompt"
	"github.com/pkg/errors"
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
