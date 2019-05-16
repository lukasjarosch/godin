package internal

import (
	"github.com/spf13/viper"
)

// Version information is injected during compilation
var (
	Version   string
	Commit string
	BuildDate string
)

// DataFromConfig loads the godin config values into the data structure
func DataFromConfig() *Data {
	data := &Data{
		Project: Project{
			RootPath: viper.GetString("project.path"),
		},
		Godin: Godin{
			Version: Version,
			Commit:  Commit,
			Build:   BuildDate,
		},
		Protobuf: Protobuf{
			Service: viper.GetString("protobuf.service"),
			Package: viper.GetString("protobuf.package"),
		},
		Service: Service{
			Name:      viper.GetString("service.name"),
			Namespace: viper.GetString("service.namespace"),
			Module:    viper.GetString("service.module"),
		},
	}

	return data
}

// Data is godin's internal data structure used to describe the current loaded project
// The data struct is also required for the templating.
type Data struct {
	Project Project
	Service Service
	Godin Godin
	Protobuf Protobuf
}

type Project struct {
	RootPath string
}

type Protobuf struct {
	Service string
	Package string
}

type Service struct {
	Name      string
	Namespace string
	Module    string
}

type Godin struct {
	Version string
	Commit string
	Build string
}
