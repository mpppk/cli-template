package option

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

type CmdConfig struct {
	*CmdRawConfig
	Out string
}

func NewRootCmdConfigFromViper() (*CmdConfig, error) {
	var conf *CmdRawConfig
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, xerrors.Errorf("failed to unmarshal config from viper: %w", err)
	}

	if err := conf.validate(); err != nil {
		return nil, xerrors.Errorf("failed to create root cmd config: %w", err)
	}

	out := conf.Out
	if conf.Out == DefaultStringValue {
		out = ""
	}
	return &CmdConfig{
		CmdRawConfig: conf,
		Out:          out,
	}, nil
}

func (c *CmdConfig) HasOut() bool {
	return c.CmdRawConfig.Out != DefaultStringValue
}

type CmdRawConfig struct {
	Norm bool
	Out string
}

func (c *CmdRawConfig) validate() error {
	if c.Out == "" {
		return fmt.Errorf("invalid --out flag value is empty")
	}
	return nil
}
