package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("netdraw")
	viper.SetDefault("install_path", "/usr/local/bin")
	viper.BindEnv("install_path")
}

// RootCmd is the root CLI command
var RootCmd = &cobra.Command{
	Use:           "netdraw",
	Short:         "Scan network and draw topology",
	SilenceUsage:  true,
	SilenceErrors: true,
}
