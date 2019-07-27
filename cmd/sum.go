package cmd

import (
	"github.com/mpppk/cli-template/lib"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "print sum of arguments",
	Long: ``,
	Args: cobra.MinimumNArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			if _, err := strconv.Atoi(arg); err != nil {
				return errors.Wrapf(err, "failed to convert args to int: ", arg)
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println(lib.SumFromString(args))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
}
