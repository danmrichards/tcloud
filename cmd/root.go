package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

var (
	// cfgFile is the path to the config file.
	cfgFile string

	// apiRegion is the name of the Tencent Cloud region to make API calls to.
	//
	// See: https://intl.cloud.tencent.com/document/product/213/15692#Region-List
	apiRegion string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "tcloud",
		Short: "A CLI for managing virtual machines in Tencent Cloud",
		Long: `A CLI for managing virtual machines in Tencent Cloud.

With this CLI you are able to create, update and delete Cloud Virtual Machine
(CVM) instances in Tencent Cloud.

A JSON config file is required to use the CLI, by default this file has to be
placed in $HOME/.tcloud.json. This can be overridden with the -config
flag. An example config file is as follows:

{
  "tencent_secret_id": "MY_SECRET_ID_GOES_HERE",
  "tencent_secret_key": "MY_SECRET_KEY_GOES_HERE"
}`,
		Args: cobra.NoArgs,
	}
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(
		&apiRegion,
		"region",
		regions.Frankfurt,
		"The Tencent Cloud API region. See: https://intl.cloud.tencent.com/document/product/213/15692#Region-List",
	)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.tcloud.json)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".tencent-cloud-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tencent-cloud-cli")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
