package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewRootCmd() (*cobra.Command, error) {
	var cmd = &cobra.Command{
		Use:   "cli-template",
		Short: "cli-template",
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-template.yaml)")
	sumCmd, err := newSumCmd()
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(sumCmd)

	selfUpdateCmd, err := newSelfUpdateCmd()
	if err != nil {
		return nil, err
	}
	cmd.AddCommand(selfUpdateCmd)

	versionCmd, err := newVersionCmd()
	cmd.AddCommand(versionCmd)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd, err := NewRootCmd()
	if err != nil {
		panic(err)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli-template")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
