package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomo",
	Short: "gomo is a cli tool to track your pomodoro sessions.",
	Long:  "gomo is a cli tool to track your pomodoro sessions. it shows you pretty timers and all ^.^",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oh no. An error while executing gomo: %s", err)
		os.Exit(1)
	}
	StartMenu()
}
