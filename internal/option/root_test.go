package option_test

import (
	"github.com/mpppk/cli-template/internal/option"
	"reflect"
	"testing"
)

func TestCmdConfig_HasOut(t *testing.T) {
	type fields struct {
		CmdRawConfig *option.CmdRawConfig
		Out          string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should return true if Out is not default value",
			fields: fields{
				CmdRawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  "test.txt",
				},
				Out: "test.txt",
			},
			want: true,
		},
		{
			name: "should return false if out is default value",
			fields: fields{
				CmdRawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  option.DefaultStringValue,
				},
				Out: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &option.CmdConfig{
				CmdRawConfig: tt.fields.CmdRawConfig,
				Out:          tt.fields.Out,
			}
			if got := c.HasOut(); got != tt.want {
				t.Errorf("HasOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRootCmdConfigFromRawConfig(t *testing.T) {
	type args struct {
		rawConfig *option.CmdRawConfig
	}
	tests := []struct {
		name string
		args args
		want *option.CmdConfig
	}{
		{
			name: "should has specified Out",
			args: args{
				rawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  "test.txt",
				},
			},
			want: &option.CmdConfig{
				CmdRawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  "test.txt",
				},
				Out: "test.txt",
			},
		},
		{
			name: "should empty Out if default value is specified",
			args: args{
				rawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  option.DefaultStringValue,
				},
			},
			want: &option.CmdConfig{
				CmdRawConfig: &option.CmdRawConfig{
					Norm: false,
					Out:  option.DefaultStringValue,
				},
				Out: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := option.NewRootCmdConfigFromRawConfig(tt.args.rawConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRootCmdConfigFromRawConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}