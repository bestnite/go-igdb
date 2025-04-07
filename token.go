package igdb

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type TwitchToken struct {
	clientID     string
	clientSecret string
	token        string
	expires      time.Time
}

func NewTwitchToken(clientID, clientSecret string) *TwitchToken {
	return &TwitchToken{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (t *TwitchToken) getToken() (string, error) {
	if t.token != "" && time.Now().Before(t.expires) {
		return t.token, nil
	}
	token, expires, err := t.loginTwitch()
	if err != nil {
		return "", fmt.Errorf("failed to login twitch: %w", err)
	}
	t.token = token
	t.expires = time.Now().Add(expires)
	return token, nil
}

func (t *TwitchToken) loginTwitch() (string, time.Duration, error) {
	baseURL, _ := url.Parse("https://id.twitch.tv/oauth2/token")
	params := url.Values{}
	params.Add("client_id", t.clientID)
	params.Add("client_secret", t.clientSecret)
	params.Add("grant_type", "client_credentials")
	baseURL.RawQuery = params.Encode()

	resp, err := request().SetHeader("User-Agent", "").Post(baseURL.String())
	if err != nil {
		return "", 0, fmt.Errorf("failed to make request: %s: %w", baseURL.String(), err)
	}

	data := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse response: %w", err)
	}
	return data.AccessToken, time.Second * time.Duration(data.ExpiresIn), nil
}
