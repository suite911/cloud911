package cmd

import (
	"fmt"
	"os"

	"github.com/suite911/cloud911/run"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if len(cfgFile) > 0 {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Search config in dirs.Config() directory with name "config" (without extension).
			viper.AddConfigPath(SysPath.Config())
			viper.SetConfigName("config")
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	})
}
