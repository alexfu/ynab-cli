// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package ynab

import (
	"ynab/internal/auth"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
)

func Budgets() ([]*budget.Summary, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	client := ynab.NewClient(token)
	return client.Budget().GetBudgets()
}
