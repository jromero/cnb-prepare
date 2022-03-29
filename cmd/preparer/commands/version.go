package commands

import (
	"fmt"

	"github.com/jromero/cnb-prepare/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number and other version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("OS / Arch:", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
