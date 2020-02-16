package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mpppk/cli-template/pkg/util"

	"github.com/mpppk/cli-template/pkg/usecase"

	"github.com/mpppk/cli-template/internal/option"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

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
					return fmt.Errorf("failed to convert args to int from %q: %w", arg, err)
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := option.NewSumCmdConfigFromViper()
			if err != nil {
				return err
			}

			util.InitializeLog(conf.Verbose)

			var result int
			if conf.Norm {
				log.Println("start L1 Norm calculation")
				r, err := usecase.CalcL1NormFromStringSlice(args)
				if err != nil {
					return fmt.Errorf("failed to calculate L1 norm: %w", err)
				}
				log.Println("finish L1 Norm calculation")
				result = r
			} else {
				log.Println("start sum calculation")
				r, err := usecase.CalcSumFromStringSlice(args)
				if err != nil {
					return fmt.Errorf("failed to calculate sum: %w", err)
				}
				log.Println("finish sum calculation")
				result = r
			}

			if conf.HasOut() {
				s := strconv.Itoa(result)
				if err := afero.WriteFile(fs, conf.Out, []byte(s), 777); err != nil {
					return fmt.Errorf("failed to write file to %s: %w", conf.Out, err)
				}
				log.Println("result is written to " + conf.Out)
			} else {
				cmd.Println(result)
			}

			return nil
		},
	}

	if err := registerSumCommandFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func registerSumCommandFlags(cmd *cobra.Command) error {
	if err := option.RegisterBoolFlag(cmd,
		&option.BoolFlag{
			Flag: &option.Flag{
				Name:  "norm",
				Usage: "Calc L1 norm instead of sum",
			},
			Value: false,
		},
	); err != nil {
		return err
	}

	if err := option.RegisterStringFlag(cmd,
		&option.StringFlag{
			Flag: &option.Flag{
				Name:  "out",
				Usage: "Output file path",
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSumCmd)
}
