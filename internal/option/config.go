package option

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag struct {
	IsPersistent bool
	Shorthand    string
	Name         string
	Usage        string
}

type StringFlag struct {
	*Flag
	Value string
}

type BoolFlag struct {
	*Flag
	Value bool
}

func getFlagSet(cmd *cobra.Command, flagConfig *Flag) (flagSet *pflag.FlagSet) {
	if flagConfig.IsPersistent {
		return cmd.PersistentFlags()
	} else {
		return cmd.Flags()
	}
}

func RegisterStringFlag(cmd *cobra.Command, flagConfig *StringFlag) error {
	flagSet := getFlagSet(cmd, flagConfig.Flag)
	if flagConfig.Shorthand == "" {
		flagSet.String(flagConfig.Name, flagConfig.Value, flagConfig.Usage)
	} else {
		flagSet.StringP(flagConfig.Name, flagConfig.Shorthand, flagConfig.Value, flagConfig.Usage)
	}

	if err := viper.BindPFlag(flagConfig.Name, flagSet.Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}

func RegisterBoolFlag(cmd *cobra.Command, flagConfig *BoolFlag) error {
	flagSet := getFlagSet(cmd, flagConfig.Flag)
	if flagConfig.Shorthand == "" {
		flagSet.Bool(flagConfig.Name, flagConfig.Value, flagConfig.Usage)
	} else {
		flagSet.BoolP(flagConfig.Name, flagConfig.Shorthand, flagConfig.Value, flagConfig.Usage)
	}

	if err := viper.BindPFlag(flagConfig.Name, flagSet.Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}
