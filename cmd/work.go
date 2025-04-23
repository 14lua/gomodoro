package cmd

import (
	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use: "work",
	Aliases: []string{"w"},
	Short: "start a work phase",
	Long: "start a 25 minute work phase",
	Run: func(cmd *cobra.Command, args []string) {
		Start(2)
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
