// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package budget

import (
	"fmt"

	"ynab/internal/auth"
	"ynab/internal/ui"
	"ynab/internal/ynab"

	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:               "list",
	Short:             "List budgets",
	PersistentPreRunE: auth.EnsureLoggedIn,
	RunE: func(cmd *cobra.Command, args []string) error {
		var budgets []*budget.Summary
		var err error
		err = ui.NewLoadingUI(func() {
			budgets, err = ynab.Budgets()
		})
		if err != nil {
			return err
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
