package cmd

import (
	"github.com/spf13/cobra"
)

var longCmd = &cobra.Command{
	Use:     "long",
	Aliases: []string{"l"},
	Short:   "start a long break",
	Long:    "start a 30 minute break",
	Run: func(cmd *cobra.Command, args []string) {
		Start(1)
	},
}

func init() {
	rootCmd.AddCommand(longCmd)
}
