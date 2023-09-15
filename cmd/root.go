package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prod",
	Short: "sku-index",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute executes the root command.
func Execute() error {
	fmt.Println("hi")
	return rootCmd.Execute()
}

func init() {}
