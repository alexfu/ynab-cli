// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"os"

	"ynab/cmd/budget"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "ynab",
	Short:        "YNAB CLI tool",
	Long:         `Manage your YNAB budget through the terminal.`,
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(budget.Cmd)
}
