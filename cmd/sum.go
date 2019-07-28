package cmd

import (
	"strconv"

	"golang.org/x/xerrors"

	"github.com/spf13/viper"

	"github.com/mpppk/cli-template/lib"
	"github.com/spf13/cobra"
)

type config struct {
	Norm bool
}

func newSumCmd() (*cobra.Command, error) {
	normFlag := &lib.BoolFlagConfig{
		Flag: &lib.Flag{
			Name:  "norm",
			Usage: "Calc L1 norm instead of sum",
		},
		Value: false,
	}

	cmd := &cobra.Command{
		Use:     "sum",
		Short:   "Print sum of arguments",
		Long:    ``,
		Args:    cobra.MinimumNArgs(2),
		Example: "cli-template sum -- -1 2  ->  1",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if _, err := strconv.Atoi(arg); err != nil {
					return xerrors.Errorf("failed to convert args to int: %s: %w", arg, err)
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var conf config
			if err := viper.Unmarshal(&conf); err != nil {
				return xerrors.Errorf("failed to unmarshal config from viper: %w", err)
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
	if err := lib.RegisterBoolFlag(cmd, normFlag); err != nil {
		return nil, err
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSumCmd)
}
