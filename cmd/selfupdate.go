package cmd

import (
	"github.com/mpppk/cli-template/internal/selfupdate"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func newSelfUpdateCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "selfupdate",
		Short: "Update cli-template",
		//Long: `Update cli-template`,
		Run: func(cmd *cobra.Command, args []string) {
			updated, err := selfupdate.Do()
			if err != nil {
				cmd.Println("Binary update failed:", err)
				return
			}
			if updated {
				cmd.Println("Current binary is the latest version", selfupdate.Version)
			} else {
				cmd.Println("Successfully updated to version", selfupdate.Version)
			}
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSelfUpdateCmd)
}
