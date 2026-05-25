// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package auth

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage ynab-cli authentication",
}
