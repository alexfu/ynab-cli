/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	"golang.org/x/term"
)

const (
	keyringService = "ynab-cli"
	keyringUser    = "default"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store your YNAB API token in the OS keyring",
	Long: `Login prompts for a YNAB Personal Access Token and stores it
  securely in your operating system's keyring (macOS Keychain, GNOME
  Keyring, or Windows Credential Manager). You only need to run this
  once per machine.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprint(os.Stderr, "YNAB Personal Access Token: ")
		tokenBytes, err := term.ReadPassword(int(os.Stdin.Fd()))

		fmt.Fprintln(os.Stderr)
		if err != nil {
			return fmt.Errorf("read token: %w", err)
		}

		token := strings.TrimSpace(string(tokenBytes))
		if token == "" {
			return fmt.Errorf("token cannot be empty")
		}

		if err := keyring.Set(keyringService, keyringUser, token); err != nil {
			return fmt.Errorf("store token in keyring: %w", err)
		}

		fmt.Fprintln(os.Stderr, "Token stored successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
