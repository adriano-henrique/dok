package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "dok",
		Short: "Fix documentation generation",
		Long:  `Dok is the best way of generating and updating documentation	for your projects`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use dok <command> for more information")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Unable to get user home directory")
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName("dok")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			isInit := false
			if len(os.Args) > 1 && os.Args[1] == "init" {
				isInit = true
			}
			if !isInit {
				fmt.Println("Config file not found. Please run 'dok init' to initialize the project.")
			}
		} else {
			fmt.Println("Error reading config file:", err)
		}
	}
}
