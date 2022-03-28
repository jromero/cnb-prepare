package commands

import (
	"fmt"
	"os"

	"github.com/jromero/cnb-prepare/pkg/preparer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "generated code example",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		preparer.Preparer(preparer.WithEnvOptions())
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
