/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of auth.",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := keyring.Get("ynab-cli", "default")
		if err == nil {
			fmt.Println("Logged in.")
		} else {
			if errors.Is(err, keyring.ErrNotFound) {
				fmt.Println("Not logged in. Run `ynab login`.")
			} else {
				return fmt.Errorf("read token from keyring: %w", err)
			}
		}
		return nil
	},
}

func init() {
	authCmd.AddCommand(statusCmd)
}
