package godin

import (
	"os"

	"fmt"

	"github.com/spf13/viper"
)

const FileName = "godin"
const FileType = "json"

// LoadConfiguration initializes viper and tries to read
// the config from the current working directory.
func LoadConfiguration() error {
	wd, _ := os.Getwd()
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)
	viper.AddConfigPath(wd)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// SaveConfiguration wraps viper.WriteConfigAs() for a cleaner API
func SaveConfiguration() error {
	return viper.WriteConfigAs(ConfigFilename())
}

// ConfigFilename returns Godin's config filename
func ConfigFilename() string {
	return fmt.Sprintf("%s.%s", FileName, FileType)
}
