package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Webhooks struct {
	request RequestFunc
}

func NewWebhooks(request RequestFunc) *Webhooks {
	return &Webhooks{
		request: request,
	}
}

type WebhookMethod string

const (
	WebhookMethodUpdate WebhookMethod = "update"
	WebhookMethodDelete WebhookMethod = "delete"
	WebhookMethodCreate WebhookMethod = "create"
)

type WebhookResponse struct {
	Id          uint64 `json:"id"`
	Url         string `json:"url"`
	Category    uint64 `json:"category"`
	SubCategory uint64 `json:"sub_category"`
	Active      bool   `json:"active"`
	ApiKey      string `json:"api_key"`
	Secret      string `json:"secret"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (a *Webhooks) Register(endpoint Name, secret, callbackUrl string, method WebhookMethod) (*WebhookResponse, error) {
	dataBody := url.Values{}
	dataBody.Set("url", callbackUrl)
	dataBody.Set("secret", secret)
	dataBody.Set("method", string(method))

	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s/webhooks/", endpoint), dataBody.Encode())

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s: %w", callbackUrl, err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil, nil
	}

	var data WebhookResponse
	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &data, fmt.Errorf("failed to activate webhook: %s: %s", callbackUrl, resp.String())
}

func (a *Webhooks) Unregister(webhookId uint64) error {
	resp, err := a.request("DELETE", fmt.Sprintf("https://api.igdb.com/v4/webhooks/%v", webhookId), "")
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return fmt.Errorf("failed to unregister webhook: %s", resp.String())
}

func (a *Webhooks) List() ([]*WebhookResponse, error) {
	resp, err := a.request("GET", "https://api.igdb.com/v4/webhooks/", "")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var data []*WebhookResponse
	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data, nil
}

func (a *Webhooks) Get(webhookId uint64) (*WebhookResponse, error) {
	resp, err := a.request("GET", fmt.Sprintf("https://api.igdb.com/v4/webhooks/%v", webhookId), "")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var data WebhookResponse
	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &data, nil
}
