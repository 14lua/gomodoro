package cmd

import (
	"github.com/spf13/cobra"
)

var shortCmd = &cobra.Command{
	Use:     "short",
	Aliases: []string{"s"},
	Short:   "start a short break",
	Long:    "start a shot five minute break",
	Run: func(cmd *cobra.Command, args []string) {
		Start(0)
	},
}

func init() {
	rootCmd.AddCommand(shortCmd)
}
