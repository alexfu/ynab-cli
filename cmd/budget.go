// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"fmt"

	"ynab/internal/auth"
	"ynab/internal/ynab"

	"github.com/spf13/cobra"
)

var budgetCmd = &cobra.Command{
	Use:               "budget",
	Short:             "List or view a budget",
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
	rootCmd.AddCommand(budgetCmd)
}
