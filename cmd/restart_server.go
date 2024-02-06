/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"palclient-cli/internal"

	"github.com/spf13/cobra"
)

// restartServerCmd represents the restartServer command
var restartServerCmd = &cobra.Command{
	Use: "restart-server",
	RunE: func(cmd *cobra.Command, args []string) error {
		notifier := internal.NewNotifier(nil, nil)

		wait, err := cmd.Flags().GetDuration("wait")
		if err != nil {
			return fmt.Errorf("getting wait duration: %w", err)
		}

		return notifier.NotifyServerRestart(wait)
	},
}

func init() {
	rootCmd.AddCommand(restartServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	restartServerCmd.Flags().DurationP("wait", "w", 0, "Waiting period before restart. Examples: 100s, 10m, 1h.")
}
