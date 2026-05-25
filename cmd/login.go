// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"fmt"

	"ynab/internal/ui"

	"github.com/spf13/cobra"
)

const (
	keyringService = "ynab-cli"
	keyringUser    = "default"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to YNAB",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ui.NewLoginUI()
		if err != nil {
			return err
		}
		fmt.Println("Log in successful!")
		return nil
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
