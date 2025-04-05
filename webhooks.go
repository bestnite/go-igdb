package igdb

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bestnite/go-igdb/endpoint"
)

func (g *Client) ActiveWebhook(endpoint endpoint.Endpoint, secret, callbackUrl string) error {
	dataBody := url.Values{}
	dataBody.Set("url", callbackUrl)
	dataBody.Set("secret", secret)
	dataBody.Set("method", "update")
	resp, err := g.Request(fmt.Sprintf("https://api.igdb.com/v4/%s/webhooks/", endpoint), dataBody.Encode())

	if err != nil {
		return fmt.Errorf("failed to make request: %s: %w", callbackUrl, err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}
	return fmt.Errorf("failed to activate webhook: %s: %s", callbackUrl, resp.String())
}
