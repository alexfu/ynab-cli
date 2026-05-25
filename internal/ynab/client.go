// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package ynab

import (
	"ynab/internal/auth"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/brunomvsouza/ynab.go/api/user"
)

func getClient() (ynab.ClientServicer, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	return ynab.NewClient(token), nil
}

func TokenValid(token string) bool {
	_, err := ynab.NewClient(token).User().GetUser()
	return err == nil
}

func Budgets() ([]*budget.Summary, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	return client.Budget().GetBudgets()
}

func GetUser() (*user.User, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	return client.User().GetUser()
}
