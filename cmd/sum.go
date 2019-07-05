package cmd

import (
	"fmt"
	"github.com/mpppk/cli-template/lib"
	"github.com/spf13/cobra"
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "print sum of arguments",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(lib.SumFromString(args))
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
}
