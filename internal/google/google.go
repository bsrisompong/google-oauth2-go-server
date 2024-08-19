package google

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bsrisompong/google-oauth-go-server/internal/config"
	"github.com/bsrisompong/google-oauth-go-server/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func InitGoogleOAuth() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectURL,
		// Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		// Endpoint:     oauth2.Endpoint{TokenURL: "https://accounts.google.com/o/oauth2/token"},
		Endpoint: google.Endpoint,
	}
}

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return googleOauthConfig.Exchange(context.TODO(), code)
}

func GetUserInfo(token *oauth2.Token) (*models.User, error) {
	client := oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(token))

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: received status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: received status code %d", resp.StatusCode)
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	return &user, nil
}
