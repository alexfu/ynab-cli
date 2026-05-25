// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear authentication with ynab-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := keyring.Delete("ynab-cli", "default")
		if err != nil {
			return fmt.Errorf("failed to logout: %w", err)
		}

		fmt.Println("Log out successful!")
		return nil
	},
}

func init() {
	authCmd.AddCommand(logoutCmd)
}
