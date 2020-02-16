package option

import (
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

// SumCmdConfig is config for sum command
type SumCmdConfig struct {
	Norm    bool
	Out     string
	Verbose bool
}

// NewSumCmdConfigFromViper generate config for sum command from viper
func NewSumCmdConfigFromViper() (*SumCmdConfig, error) {
	var conf SumCmdConfig
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal config from viper: %w", err)
	}

	if err := conf.validate(); err != nil {
		return nil, xerrors.Errorf("failed to create sum cmd config: %w", err)
	}
	return &conf, nil
}

// HasOut returns whether or not config has Out property
func (c *SumCmdConfig) HasOut() bool {
	return c.Out != ""
}

func (c *SumCmdConfig) validate() error {
	return nil
}
