package cmd

import (
	"github.com/mpppk/cli-template/internal/option"
	"github.com/mpppk/cli-template/pkg/sum"
	"github.com/mpppk/cli-template/pkg/util"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

type config struct {
	Norm bool
}

func newSumCmd() (*cobra.Command, error) {
	normFlag := &option.BoolFlag{
		Flag: &option.Flag{
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

			numbers, err := util.ConvertStringSliceToIntSlice(args)
			if err != nil {
				return err
			}

			var res int
			if conf.Norm {
				res = sum.L1Norm(numbers)
			} else {
				res = sum.Sum(numbers)
			}
			cmd.Println(res)
			return nil
		},
	}
	if err := option.RegisterBoolFlag(cmd, normFlag); err != nil {
		return nil, err
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSumCmd)
}
