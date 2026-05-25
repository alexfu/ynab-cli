// Copyright © 2026 Alex Fu <alexfu@fastmail.com>

package auth

import (
	"github.com/zalando/go-keyring"
)

const (
	keyringService = "com.alexfu.ynab-cli"
	keyringUser    = "default"
)

func SaveToken(token string) error {
	return keyring.Set(keyringService, keyringUser, token)
}

func Logout() error {
	return keyring.Delete(keyringService, keyringUser)
}

func HasToken() bool {
	_, err := GetToken()
	return err == nil
}

func GetToken() (string, error) {
	return keyring.Get(keyringService, keyringUser)
}
