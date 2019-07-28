package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mpppk/cli-template/cmd"
)

func TestSum(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "sum -- -1 2", want: "1\n"},
		{command: "sum --norm -- -1 2", want: "3\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		rootCmd, err := cmd.NewRootCmd()
		if err != nil {
			t.Errorf("failed to create rootCmd: %s", err)
		}
		rootCmd.SetOut(buf)
		cmdArgs := strings.Split(c.command, " ")
		rootCmd.SetArgs(cmdArgs)
		if err := rootCmd.Execute(); err != nil {
			t.Errorf("failed to execute rootCmd: %s", err)
		}

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%q, get:%q", c.want, get)
		}
	}
}
