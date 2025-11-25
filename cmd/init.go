package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CONFIG_FILE_NAME = "dok.yaml"

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new dok project",
	Long:  `Initialize a new dok project`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("author", "YOUR NAME")
		viper.Set("license", "apache")

		if err := viper.SafeWriteConfigAs(CONFIG_FILE_NAME); err != nil {
			if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
				fmt.Println("Config file already exists")
				return
			}
			fmt.Println("Error creating config file:", err)
			return
		}
		fmt.Println("Created dok.yaml")
	},
}
