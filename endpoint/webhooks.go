package endpoint

import (
	"fmt"
	"net/http"
	"net/url"
)

type Webhooks struct{ BaseEndpoint }

func (a *Webhooks) Register(endpoint EndpointName, secret, callbackUrl string) error {
	dataBody := url.Values{}
	dataBody.Set("url", callbackUrl)
	dataBody.Set("secret", secret)
	dataBody.Set("method", "update")
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s/webhooks/", endpoint), dataBody.Encode())

	if err != nil {
		return fmt.Errorf("failed to make request: %s: %w", callbackUrl, err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}
	return fmt.Errorf("failed to activate webhook: %s: %s", callbackUrl, resp.String())
}
