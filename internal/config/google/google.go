package google

import (
	"context"

	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func InitGoogleOAuth() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectURL,
		// Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		// Endpoint:     oauth2.Endpoint{TokenURL: "https://accounts.google.com/o/oauth2/token"},
		Endpoint: google.Endpoint,
	}
}

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return googleOauthConfig.Exchange(context.TODO(), code)
}
