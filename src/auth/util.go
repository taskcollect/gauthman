package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var TC_API_SCOPES = []string{"https://www.googleapis.com/auth/userinfo.email"}

type OAuth2Secrets struct {
	ClientID     string
	ClientSecret string
}

func GetOAuth2Config(secrets *OAuth2Secrets, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		RedirectURL:  "postmessage",
		Scopes:       TC_API_SCOPES,
		Endpoint:     google.Endpoint,
	}
}
