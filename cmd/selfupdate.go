package cmd

import (
	"github.com/mpppk/cli-template/lib"
	"github.com/spf13/cobra"
	"log"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "selfupdate",
	Short: "update cli-template",
	//Long: `Update cli-template`,
	Run: func(cmd *cobra.Command, args []string) {
		updated, err := lib.DoSelfUpdate()
		if err != nil {
			log.Println("Binary update failed:", err)
			return
		}
		if updated {
			log.Println("Current binary is the latest version", lib.Version)
		} else {
			log.Println("Successfully updated to version", lib.Version)
		}
	},
}

func init() {
	rootCmd.AddCommand(selfUpdateCmd)
}
