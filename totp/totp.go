package totp

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func NewKey(name, issuer string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: name,
	})
	if err != nil {
		return nil, fmt.Errorf("totp.Generate error: %w", err)
	}

	return key, nil
}

func NewKeyFromUrl(url string) (*otp.Key, error) {
	return otp.NewKeyFromURL(url)
}
