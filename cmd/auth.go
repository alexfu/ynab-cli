// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage ynab-cli authentication",
}

func init() {
	rootCmd.AddCommand(authCmd)
}
