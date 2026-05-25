// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package auth

import (
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

const (
	keyringService = "com.alexfu.ynab-cli"
	keyringUser    = "default"
)

func EnsureLoggedIn(cmd *cobra.Command, args []string) error {
	if !LoggedIn() {
		err := NewLoginUI()
		if err != nil {
			return err
		}
	}
	return nil
}

func Login(token string) error {
	return keyring.Set(keyringService, keyringUser, token)
}

func Logout() error {
	return keyring.Delete(keyringService, keyringUser)
}

func LoggedIn() bool {
	_, err := GetToken()
	return err == nil
}

func GetToken() (string, error) {
	return keyring.Get(keyringService, keyringUser)
}
