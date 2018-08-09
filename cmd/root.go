package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amy911/amy911/syspath"

	"github.com/amy911/srv911/run"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Set this in your `init()` function somewhere
	VendorName = "amy911"

	// Set this in your `init()` function somewhere
	ApplicationName = "srv911"

	// Set this in your `init()` function somewhere
	DefaultConfType = "yaml"

	// Set this in your `init()` function somewhere
	OverrideRoot func(*cobra.Command, []string)
)

var (
	// Not usable in `init()` functions
	SysPath *syspath.SysPath

	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   ApplicationName,
	Short: "An application server over HTTP and HTTPS",
	Long: `An application server over HTTP and HTTPS based on [srv911](https://github.com/amy911/srv911)`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case OverrideRoot == nil:
			run.Listen(cmd, args)
		default:
			OverrideRoot(cmd, args)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	SysPath = syspath.New(VendorName, ApplicationName)
	if len(cfgFile) > 0 {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in dirs.Config() directory with name "ApplicationName" (without extension).
		viper.AddConfigPath(SysPath.Config())
		viper.SetConfigName(ApplicationName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
