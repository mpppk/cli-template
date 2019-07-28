package cmd

import (
	"github.com/mpppk/cli-template/internal/selfupdate"
	"github.com/spf13/cobra"
)

func newVersionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		//Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(selfupdate.Version)
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newVersionCmd)
}
