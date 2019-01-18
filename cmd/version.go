package cmd

import (
	"fmt"
	"github.com/mpppk/cli-template/lib"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(lib.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
