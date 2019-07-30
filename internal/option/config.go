// Package option provides utilities of option handling
package option

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flag represents flag which can be specified
type Flag struct {
	IsPersistent bool
	IsRequired   bool
	IsDirName    bool
	IsFileName   bool
	Shorthand    string
	Name         string
	Usage        string
}

// StringFlag represents flag which can be specified as string
type StringFlag struct {
	*Flag
	Value string
}

// BoolFlag represents flag which can be specified as bool
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

// RegisterStringFlag register string flag to provided cmd and viper
func RegisterStringFlag(cmd *cobra.Command, flagConfig *StringFlag) error {
	flagSet := getFlagSet(cmd, flagConfig.Flag)
	if flagConfig.Shorthand == "" {
		flagSet.String(flagConfig.Name, flagConfig.Value, flagConfig.Usage)
	} else {
		flagSet.StringP(flagConfig.Name, flagConfig.Shorthand, flagConfig.Value, flagConfig.Usage)
	}

	if err := markAttributes(cmd, flagConfig.Flag); err != nil {
		return err
	}

	if err := viper.BindPFlag(flagConfig.Name, flagSet.Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}

// RegisterBoolFlag register bool flag to provided cmd and viper
func RegisterBoolFlag(cmd *cobra.Command, flagConfig *BoolFlag) error {
	flagSet := getFlagSet(cmd, flagConfig.Flag)
	if flagConfig.Shorthand == "" {
		flagSet.Bool(flagConfig.Name, flagConfig.Value, flagConfig.Usage)
	} else {
		flagSet.BoolP(flagConfig.Name, flagConfig.Shorthand, flagConfig.Value, flagConfig.Usage)
	}

	if err := markAttributes(cmd, flagConfig.Flag); err != nil {
		return err
	}

	if err := viper.BindPFlag(flagConfig.Name, flagSet.Lookup(flagConfig.Name)); err != nil {
		return err
	}
	return nil
}

func markAttributes(cmd *cobra.Command, flagConfig *Flag) error {
	if err := markAsFileName(cmd, flagConfig); err != nil {
		return err
	}
	if err := markAsDirName(cmd, flagConfig); err != nil {
		return err
	}
	if err := markAsRequired(cmd, flagConfig); err != nil {
		return err
	}
	return nil
}

func markAsFileName(cmd *cobra.Command, flagConfig *Flag) error {
	if flagConfig.IsFileName {
		if flagConfig.IsPersistent {
			if err := cmd.MarkPersistentFlagFilename(flagConfig.Name); err != nil {
				return err
			}
		} else {
			if err := cmd.MarkFlagFilename(flagConfig.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

func markAsDirName(cmd *cobra.Command, flagConfig *Flag) error {
	if flagConfig.IsDirName {
		if flagConfig.IsPersistent {
			if err := cmd.MarkPersistentFlagDirname(flagConfig.Name); err != nil {
				return err
			}
		} else {
			if err := cmd.MarkFlagDirname(flagConfig.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

func markAsRequired(cmd *cobra.Command, flagConfig *Flag) error {
	if flagConfig.IsRequired {
		if flagConfig.IsPersistent {
			if err := cmd.MarkPersistentFlagRequired(flagConfig.Name); err != nil {
				return err
			}
		} else {
			if err := cmd.MarkFlagRequired(flagConfig.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
