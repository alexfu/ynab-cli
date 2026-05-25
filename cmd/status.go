// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"fmt"

	"ynab/internal/auth"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current status of auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		if auth.LoggedIn() {
			fmt.Println("Logged in.")
		} else {
			fmt.Println("Not logged in. Run `ynab auth login` to log in.")
		}
		return nil
	},
}

func init() {
	authCmd.AddCommand(statusCmd)
}
