// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package budget

import (
	"fmt"

	"ynab/internal/auth"
	"ynab/internal/ynab"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:               "list",
	Short:             "List budgets",
	PersistentPreRunE: auth.EnsureLoggedIn,
	RunE: func(cmd *cobra.Command, args []string) error {
		budgets, err := ynab.Budgets()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for _, budget := range budgets {
			fmt.Printf("%v\t%v\n", budget.ID, budget.Name)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(listCmd)
}
