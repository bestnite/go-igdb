package igdb

import (
	"fmt"
	"github/bestnite/go-igdb/constant"
	"net/http"
	"net/url"
)

// ActiveWebhook activates a webhook for a specific endpoint.
//
// https://api-docs.igdb.com/#webhooks
func (g *igdb) ActiveWebhook(endpoint, secret, callbackUrl string) error {
	t, err := g.token.getToken()
	if err != nil {
		return fmt.Errorf("failed to get Twitch token: %w", err)
	}
	dataBody := url.Values{}
	dataBody.Set("url", callbackUrl)
	dataBody.Set("secret", secret)
	dataBody.Set("method", "update")
	resp, err := request().SetBody(dataBody.Encode()).SetHeaders(map[string]string{
		"Client-ID":     g.clientID,
		"Authorization": "Bearer " + t,
		"User-Agent":    "",
		"Content-Type":  "application/x-www-form-urlencoded",
	}).Post(fmt.Sprintf(constant.IGDBWebhookURL, endpoint))

	if err != nil {
		return fmt.Errorf("failed to make request: %s: %w", callbackUrl, err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}
	return fmt.Errorf("failed to activate webhook: %s: %s", callbackUrl, resp.String())
}
