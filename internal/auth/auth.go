// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package auth

import "github.com/zalando/go-keyring"

const (
	keyringService = "com.alexfu.ynab-cli"
	keyringUser    = "default"
)

func Login(token string) error {
	return keyring.Set(keyringService, keyringUser, token)
}

func Logout() error {
	return keyring.Delete(keyringService, keyringUser)
}

func LoggedIn() bool {
	_, err := keyring.Get(keyringService, keyringUser)
	return err == nil
}
