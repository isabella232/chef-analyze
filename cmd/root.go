package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/chef/chef-analyze/pkg/config"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "chef-analyze",
		Short: "Analyze your Chef inventory",
		Long: `Analyze your Chef inventory by generating reports to understand the effort
to upgrade to the latest version of the Chef tools, automatically fixing
violations and/or deprecations, and generating Effortless packages.
`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Chef config file (default: $HOME/.chef/credentials)")
	rootCmd.PersistentFlags().StringP("client_name", "n", "", "Chef Infra Server API client username")
	rootCmd.PersistentFlags().StringP("client_key", "k", "", "Chef Infra Server API client key")
	rootCmd.PersistentFlags().StringP("chef_server_url", "s", "", "Chef Infra Server URL")
	rootCmd.PersistentFlags().StringP("profile", "p", "", "Chef Infra Server URL")
	viper.BindPFlag("client_name", rootCmd.PersistentFlags().Lookup("client_name"))
	viper.BindPFlag("client_key", rootCmd.PersistentFlags().Lookup("client_key"))
	viper.BindPFlag("chef_server_url", rootCmd.PersistentFlags().Lookup("chef_server_url"))

	// adds the report command from 'cmd/report.go'
	rootCmd.AddCommand(reportCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find the config and pass it to viper
		configFile, err := config.FindConfigFile()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			fmt.Println(rootCmd.UsageString())
			os.Exit(-1)
		}
		viper.SetConfigFile(configFile)
	}

	viper.SetConfigType("toml")
	viper.AutomaticEnv()
}
