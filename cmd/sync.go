package cmd

import (
	"github.com/lukasjarosch/godin/internal"
	"github.com/lukasjarosch/godin/internal/project"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
	"path"
)

func init() {
	rootCmd.AddCommand(syncCommand)
}

// rootCmd represents the base command when called without any subcommands
var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync the generated files with the service interface",
	Run:   syncCmd,
}

func syncCmd(cmd *cobra.Command, args []string) {
	if err := project.HasConfig(); err != nil {
		logrus.Fatal("project not initialized")
	}

	// setup the template data
	data := internal.DataFromConfig()

	fileNames := []string{
		"internal/service/" + data.Service.Name + "/implementation.go",
		"internal/service/middleware/logging.go",
	}

	var files []*types.File
	for _, fileName := range fileNames {
		f, err := astra.ParseFile(path.Join(data.Project.RootPath, fileName))
		if err != nil {
			logrus.Fatal(err)
		}
		files = append(files, f)
	}

	parsed, err := astra.MergeFiles(files)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("parsed all files: %s", parsed.Name)

	// for each file to be modified
	// parse with astra and check which functions already exist
	// render partial templates of those who are missing
	// append those partials to the file
	// save
}
