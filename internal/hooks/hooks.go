package hooks

import (
	"ynab/internal/auth"
	"ynab/internal/ui"

	"github.com/spf13/cobra"
)

func EnsureLoggedIn(cmd *cobra.Command, args []string) error {
	if !auth.LoggedIn() {
		err := ui.NewLoginUI()
		if err != nil {
			return err
		}
	}
	return nil
}
