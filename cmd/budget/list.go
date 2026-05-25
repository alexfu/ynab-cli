// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package budget

import (
	"fmt"

	"ynab/internal/hooks"
	"ynab/internal/ui"
	"ynab/internal/utils"
	"ynab/internal/ynab"

	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:               "list",
	Short:             "List budgets",
	PersistentPreRunE: hooks.EnsureLoggedIn,
	RunE: func(cmd *cobra.Command, args []string) error {
		var budgets []*budget.Summary
		var err error
		err = ui.NewLoadingUI(func() {
			budgets, err = ynab.Budgets()
		})
		if err != nil {
			return err
		}

		jsonOutput, _ := cmd.Flags().GetBool("json")
		if jsonOutput {
			jsonStr, _ := utils.ToJSONString(budgets)
			fmt.Println(jsonStr)
		} else {
			for _, budget := range budgets {
				fmt.Printf("%v\t%v\n", budget.ID, budget.Name)
			}
		}
		return nil
	},
}

func init() {
	listCmd.Flags().Bool("json", false, "Output as JSON")
	Cmd.AddCommand(listCmd)
}
