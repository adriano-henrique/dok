package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dok",
	Short: "Fix documentation generation",
	Long:  `Dok is the best way of generating and updating documentation	for your projects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use dok <command> for more information")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
