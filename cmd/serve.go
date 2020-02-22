package cmd

import (
	"github.com/mpppk/cli-template/handler"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

func newServeCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			e := handler.NewServer()
			e.Logger.Fatal(e.Start(":1323"))
			return nil
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newServeCmd)
}
