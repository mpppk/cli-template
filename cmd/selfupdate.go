package cmd

import (
	"github.com/mpppk/cli-template/lib"
	"github.com/spf13/cobra"
)

func newSelfUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "selfupdate",
		Short: "Update cli-template",
		//Long: `Update cli-template`,
		Run: func(cmd *cobra.Command, args []string) {
			updated, err := lib.DoSelfUpdate()
			if err != nil {
				cmd.Println("Binary update failed:", err)
				return
			}
			if updated {
				cmd.Println("Current binary is the latest version", lib.Version)
			} else {
				cmd.Println("Successfully updated to version", lib.Version)
			}
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSelfUpdateCmd)
}
