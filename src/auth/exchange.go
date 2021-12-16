package auth

import (
	"context"

	"golang.org/x/oauth2"
)

func ExchangeCodeForTokenPair(auth_code string, secrets *OAuth2Secrets) (*oauth2.Token, error) {
	// make a google oauth2 oa2cfg for offline access
	oa2cfg := GetOAuth2Config(secrets, TC_API_SCOPES)

	// exchange auth code for tokens
	token, err := oa2cfg.Exchange(context.Background(), auth_code, oauth2.AccessTypeOffline)

	if err != nil {
		return nil, err
	}

	return token, nil
}
