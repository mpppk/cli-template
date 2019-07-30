package option

import (
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

type CmdConfig struct {
	Toggle bool
}

func NewRootCmdConfigFromViper() (*CmdConfig, error) {
	rawConfig, err := newCmdRawConfig()
	return newCmdConfigFromRawConfig(rawConfig), err
}

func newCmdRawConfig() (*CmdRawConfig, error) {
	var conf CmdRawConfig
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal config from viper: %w", err)
	}

	if err := conf.validate(); err != nil {
		return nil, xerrors.Errorf("failed to create root cmd config: %w", err)
	}
	return &conf, nil
}

func newCmdConfigFromRawConfig(rawConfig *CmdRawConfig) *CmdConfig {
	return &CmdConfig{
		Toggle: rawConfig.Toggle,
	}
}

type CmdRawConfig struct {
	SumCmdConfig `mapstructure:",squash"`
	Toggle       bool
}

func (c *CmdRawConfig) validate() error {
	if err := c.SumCmdConfig.validate(); err != nil {
		return xerrors.Errorf("invalid config parameter is given to SumCmdConfig: %w", err)
	}
	return nil
}
