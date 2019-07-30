package cmd

import (
	"strconv"

	"github.com/mpppk/cli-template/internal/option"
	"github.com/mpppk/cli-template/pkg/sum"
	"github.com/mpppk/cli-template/pkg/util"
	"github.com/spf13/afero"

	"golang.org/x/xerrors"

	"github.com/spf13/cobra"
)

func newNormFlag() *option.BoolFlag {
	return &option.BoolFlag{
		Flag: &option.Flag{
			Name:  "norm",
			Usage: "Calc L1 norm instead of sum",
		},
		Value: false,
	}
}

func newOutFlag() *option.StringFlag {
	return &option.StringFlag{
		Flag: &option.Flag{
			Name:  "out",
			Usage: "Output file path",
		},
		Value: option.DefaultStringValue,
	}
}

func newSumCmd(fs afero.Fs) (*cobra.Command, error) {
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
			conf, err := option.NewSumCmdConfigFromViper()
			if err != nil {
				return err
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

			if conf.HasOut() {
				s := strconv.Itoa(res)
				if err := afero.WriteFile(fs, conf.Out, []byte(s), 777); err != nil {
					return xerrors.Errorf("failed to write file to %s: %w", conf.Out, err)
				}
			} else {
				cmd.Println(res)
			}

			return nil
		},
	}
	if err := option.RegisterBoolFlag(cmd, newNormFlag()); err != nil {
		return nil, err
	}
	if err := option.RegisterStringFlag(cmd, newOutFlag()); err != nil {
		return nil, err
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSumCmd)
}
