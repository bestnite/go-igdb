package igdb

import (
	"fmt"

	"git.nite07.com/nite/flaresolverr"
	"github.com/go-resty/resty/v2"
)

type igdb struct {
	clientID     string
	token        *twitchToken
	flaresolverr *flaresolverr.Flaresolverr
}

func New(clientID, clientSecret string) *igdb {
	return &igdb{
		clientID:     clientID,
		token:        NewTwitchToken(clientID, clientSecret),
		flaresolverr: nil,
	}
}

func NewWithFlaresolverr(clientID, clientSecret string, f *flaresolverr.Flaresolverr) *igdb {
	return &igdb{
		clientID:     clientID,
		token:        NewTwitchToken(clientID, clientSecret),
		flaresolverr: f,
	}
}

func (g *igdb) Request(URL string, dataBody any) (*resty.Response, error) {
	t, err := g.token.getToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get Twitch token: %w", err)
	}

	resp, err := request().SetBody(dataBody).SetHeaders(map[string]string{
		"Client-ID":     g.clientID,
		"Authorization": "Bearer " + t,
		"User-Agent":    "",
		"Content-Type":  "text/plain",
	}).Post(URL)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %s: %w", URL, err)
	}
	return resp, nil
}

func (g *igdb) getFlaresolverr() (*flaresolverr.Flaresolverr, error) {
	if g.flaresolverr == nil {
		return nil, fmt.Errorf("flaresolverr is not initialized")
	}
	return g.flaresolverr, nil
}
