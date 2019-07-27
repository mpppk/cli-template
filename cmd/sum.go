package cmd

import (
	"strconv"

	"github.com/spf13/viper"

	"github.com/mpppk/cli-template/lib"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var normFlag = &lib.BoolFlagConfig{
	Name:  "norm",
	Value: false,
	Usage: "calc L1 norm instead of sum",
}

type config struct {
	Norm bool
}

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "print sum of arguments",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			if _, err := strconv.Atoi(arg); err != nil {
				return errors.Wrapf(err, "failed to convert args to int: %s", arg)
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var conf config
		if err := viper.Unmarshal(&conf); err != nil {
			return errors.Wrap(err, "failed to unmarshal config from viper")
		}

		numbers, err := lib.ConvertStringSliceToIntSlice(args)
		if err != nil {
			return err
		}

		var res int
		if conf.Norm {
			res = lib.L1Norm(numbers)
		} else {
			res = lib.Sum(numbers)
		}
		cmd.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
	if err := lib.RegisterBoolFlag(sumCmd, normFlag); err != nil {
		panic(err)
	}
}
