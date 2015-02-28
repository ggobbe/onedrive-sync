package api

import (
	"fmt"
	"net/http"

	"github.com/ggobbe/onedrive-sync/config"
	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:    "000000004C140112",
	Scopes:      []string{"onedrive.readonly", "wl.signin"},
	RedirectURL: "https://login.live.com/oauth20_desktop.srf",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://login.live.com/oauth20_authorize.srf",
		TokenURL: "https://login.live.com/oauth20_token.srf",
	},
}

func getClient() (*http.Client, error) {
	tok, err := getToken()
	if err != nil {
		return nil, err
	}
	return conf.Client(oauth2.NoContext, tok), nil
}

func askCode() (string, error) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Please visit the URL for the auth dialog:\n%v\n", url)
	fmt.Print("and copy/paster here the code from the callback URL: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return "", err
	}
	return code, nil
}

var token *oauth2.Token

func getToken() (*oauth2.Token, error) {
	if token != nil {
		return token, nil
	}
	code, err := askCode()
	if err != nil {
		return nil, err
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	err = config.SaveToken(tok)
	return tok, err
}
