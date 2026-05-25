// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"errors"
	"fmt"

	"ynab/internal/auth"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of YNAB",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := auth.Logout()
		if err != nil {
			if errors.Is(err, keyring.ErrNotFound) {
				fmt.Println("Already logged out!")
				return nil
			} else {
				return fmt.Errorf("failed to logout: %w", err)
			}
		}

		fmt.Println("Log out successful!")
		return nil
	},
}

func init() {
	authCmd.AddCommand(logoutCmd)
}
