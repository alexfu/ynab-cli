// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"ynab/internal/auth"
	"ynab/internal/ynab"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current status of auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		if auth.LoggedIn() {
			user, err := ynab.GetUser()
			if err != nil {
				return err
			}
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			fmt.Fprintf(w, "Status:\tLogged In\nUser ID:\t%v\n", user.ID)
			w.Flush()
		} else {
			fmt.Println("Not logged in. Run `ynab auth login` to log in.")
		}
		return nil
	},
}

func init() {
	authCmd.AddCommand(statusCmd)
}
