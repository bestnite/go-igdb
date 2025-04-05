package igdb

import (
	"fmt"
	"strings"

	"github.com/bestnite/go-flaresolverr"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	clientID     string
	token        *twitchToken
	flaresolverr *flaresolverr.Flaresolverr
	limiter      *rateLimiter
}

func New(clientID, clientSecret string) *Client {
	return &Client{
		clientID:     clientID,
		limiter:      newRateLimiter(4),
		token:        NewTwitchToken(clientID, clientSecret),
		flaresolverr: nil,
	}
}

func NewWithFlaresolverr(clientID, clientSecret string, f *flaresolverr.Flaresolverr) *Client {
	return &Client{
		clientID:     clientID,
		limiter:      newRateLimiter(4),
		token:        NewTwitchToken(clientID, clientSecret),
		flaresolverr: f,
	}
}

func (g *Client) Request(URL string, dataBody any) (*resty.Response, error) {
	g.limiter.wait()

	t, err := g.token.getToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get twitch token: %w", err)
	}

	resp, err := request().SetBody(dataBody).SetHeaders(map[string]string{
		"Client-ID":     g.clientID,
		"Authorization": "Bearer " + t,
		"User-Agent":    "",
		"Content-Type":  "text/plain",
	}).Post(URL)

	if err != nil {
		return nil, fmt.Errorf("failed to request: %s: %w", URL, err)
	}
	return resp, nil
}

func (g *Client) getFlaresolverr() (*flaresolverr.Flaresolverr, error) {
	if g.flaresolverr == nil {
		return nil, fmt.Errorf("flaresolverr is not initialized")
	}
	return g.flaresolverr, nil
}

func GetItemByID[T any](id uint64, f func(string) ([]*T, error)) (*T, error) {
	query := fmt.Sprintf("where id = %d; fields *;", id)
	items, err := f(query)
	if err != nil {
		return nil, err
	}
	return items[0], nil
}

func GetItemsByIDs[T any](ids []uint64, f func(string) ([]*T, error)) ([]*T, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return f(idStr)
}

func GetItemsPagniated[T any](offset, limit int, f func(string) ([]*T, error)) ([]*T, error) {
	query := fmt.Sprintf("offset %d; limit %d; f *; sort id asc;", offset, limit)
	return f(query)
}

func GetItemsLength[T any](f func(string) ([]*T, error)) (int, error) {
	query := "fields id; sort id desc; limit 1;"
	items, err := f(query)
	if err != nil {
		return 0, err
	}
	return len(items), nil
}
