package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test command",
	Long:  `Test command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use test <command> for more information")
	},
}
