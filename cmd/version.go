package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	format := "Godin - The awesome go-kit project manager: v%s (%s), built: %s\n"
	fmt.Printf(format, Version, Commit, BuildDate)
}
