package cmd

import (
	"bytes"
	"strings"
	"testing"
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
		cmd, err := NewRootCmd()
		if err != nil {
			t.Errorf("failed to create root cmd: %s", err)
		}
		cmd.SetOut(buf)
		cmdArgs := strings.Split(c.command, " ")
		cmd.SetArgs(cmdArgs)
		if err := cmd.Execute(); err != nil {
			t.Errorf("failed to execute cmd: %s", err)
		}

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%q, get:%q", c.want, get)
		}
	}
}
