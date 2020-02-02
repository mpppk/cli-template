package cmd

import (
	"strconv"

	"github.com/mpppk/cli-template/pkg/usecase"

	"github.com/mpppk/cli-template/internal/option"
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

			var result int
			if conf.Norm {
				r, err := usecase.CalcL1NormFromStringSlice(args)
				if err != nil {
					return xerrors.Errorf("failed to calculate L1 norm: %w", err)
				}
				result = r
			} else {
				r, err := usecase.CalcSumFromStringSlice(args)
				if err != nil {
					return xerrors.Errorf("failed to calculate sum: %w", err)
				}
				result = r
			}

			if conf.HasOut() {
				s := strconv.Itoa(result)
				if err := afero.WriteFile(fs, conf.Out, []byte(s), 777); err != nil {
					return xerrors.Errorf("failed to write file to %s: %w", conf.Out, err)
				}
			} else {
				cmd.Println(result)
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
