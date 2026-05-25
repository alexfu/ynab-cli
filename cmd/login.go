// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

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
	Short: "Authenticate with ynab-cli",
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
	authCmd.AddCommand(loginCmd)
}
