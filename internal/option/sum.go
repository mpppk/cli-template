package option

import "fmt"

// SumCmdConfig is config for sum command
type SumCmdConfig struct {
	Norm bool
	Out  string
}

// NewSumCmdConfigFromViper generate config for sum command from viper
func NewSumCmdConfigFromViper() (*SumCmdConfig, error) {
	rawConfig, err := newCmdRawConfig()
	return newSumCmdConfigFromRawConfig(rawConfig), err
}

func newSumCmdConfigFromRawConfig(rawConfig *CmdRawConfig) *SumCmdConfig {
	out := rawConfig.Out
	if rawConfig.Out == DefaultStringValue {
		out = ""
	}
	return &SumCmdConfig{
		Norm: rawConfig.Norm,
		Out:  out,
	}
}

// HasOut returns whether or not config has Out property
func (c *SumCmdConfig) HasOut() bool {
	return c.Out != ""
}

func (c *SumCmdConfig) validate() error {
	if c.Out == "" {
		return fmt.Errorf("invalid --out flag value is empty")
	}
	return nil
}
