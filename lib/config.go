package lib

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type StringPFlagConfig struct {
	Name      string
	Shorthand string
	Value     string
	Usage     string
}

func RegisterStringPFlag(cmd *cobra.Command, flagConfig *StringPFlagConfig) error {
	cmd.Flags().StringP(flagConfig.Name, flagConfig.Shorthand, flagConfig.Value, flagConfig.Usage)
	if err := viper.BindPFlag(flagConfig.Name, cmd.Flags().Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}

type BoolFlagConfig struct {
	Name  string
	Value bool
	Usage string
}

func RegisterBoolFlag(cmd *cobra.Command, flagConfig *BoolFlagConfig) error {
	cmd.Flags().Bool(flagConfig.Name, flagConfig.Value, flagConfig.Usage)
	if err := viper.BindPFlag(flagConfig.Name, cmd.Flags().Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}
