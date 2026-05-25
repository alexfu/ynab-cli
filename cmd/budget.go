// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package cmd

import (
	"errors"
	"fmt"

	"ynab/internal/auth"
	"ynab/internal/ynab"

	"github.com/spf13/cobra"
)

var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "List or view a budget",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !auth.LoggedIn() {
			return errors.New("not logged in")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		budgets, _ := ynab.Budgets()
		fmt.Println(budgets)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(budgetCmd)
}
