package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lukasjarosch/godin/internal"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

// rootCmd represents the base command when called without any subcommands
var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show the Godin version information",
	Run:   versionCmd,
}

func versionCmd(cmd *cobra.Command, args []string) {
	format := "Godin - go-kit project manager\nv%s (%s), built: %s\n"
	fmt.Printf(format, internal.Version, internal.Commit, internal.BuildDate)
}
