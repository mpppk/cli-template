package option

import "fmt"

type SumCmdConfig struct {
	Norm bool
	Out  string
}

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

func (c *SumCmdConfig) HasOut() bool {
	return c.Out != ""
}

func (c *SumCmdConfig) validate() error {
	if c.Out == "" {
		return fmt.Errorf("invalid --out flag value is empty")
	}
	return nil
}
