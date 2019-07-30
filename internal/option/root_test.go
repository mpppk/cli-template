package option_test

import (
	"reflect"
	"testing"

	"github.com/mpppk/cli-template/internal/option"
)

func Test_newCmdConfigFromRawConfig(t *testing.T) {
	type args struct {
		rawConfig *option.CmdRawConfig
	}
	tests := []struct {
		name string
		args args
		want *option.CmdConfig
	}{
		{
			name: "Toggle property should have false if CmdRawConfig has false",
			args: args{
				rawConfig: &option.CmdRawConfig{
					SumCmdConfig: option.SumCmdConfig{
						Norm: false,
						Out:  "",
					},
					Toggle: false,
				},
			},
			want: &option.CmdConfig{
				Toggle: false,
			},
		},
		{
			name: "Toggle property should have true if CmdRawConfig has true",
			args: args{
				rawConfig: &option.CmdRawConfig{
					SumCmdConfig: option.SumCmdConfig{
						Norm: false,
						Out:  "",
					},
					Toggle: true,
				},
			},
			want: &option.CmdConfig{
				Toggle: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := option.NewCmdConfigFromRawConfig(tt.args.rawConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCmdConfigFromRawConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
