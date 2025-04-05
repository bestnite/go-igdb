package igdb

import (
	"fmt"
	"strings"

	"github.com/bestnite/go-flaresolverr"
	"github.com/bestnite/go-igdb/endpoint"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	clientID     string
	token        *twitchToken
	flaresolverr *flaresolverr.Flaresolverr
	limiter      *rateLimiter

	EntityEndpoints map[endpoint.EndpointName]endpoint.EntityEndpoint

	AgeRatingCategories            *endpoint.AgeRatingCategories
	AgeRatingContentDescriptions   *endpoint.AgeRatingContentDescriptions
	AgeRatingContentDescriptionsV2 *endpoint.AgeRatingContentDescriptionsV2
	AgeRatingOrganizations         *endpoint.AgeRatingOrganizations
	AgeRatings                     *endpoint.AgeRatings
	AlternativeNames               *endpoint.AlternativeNames
	Artworks                       *endpoint.Artworks
	CharacterGenders               *endpoint.CharacterGenders
	CharacterMugShots              *endpoint.CharacterMugShots
	Characters                     *endpoint.Characters
	CharacterSpecies               *endpoint.CharacterSpecies
	CollectionMemberships          *endpoint.CollectionMemberships
	CollectionMembershipTypes      *endpoint.CollectionMembershipTypes
	CollectionRelations            *endpoint.CollectionRelations
	CollectionRelationTypes        *endpoint.CollectionRelationTypes
	Collections                    *endpoint.Collections
	CollectionTypes                *endpoint.CollectionTypes
	Companies                      *endpoint.Companies
	CompanyLogos                   *endpoint.CompanyLogos
	CompanyStatuses                *endpoint.CompanyStatuses
	CompanyWebsites                *endpoint.CompanyWebsites
	Covers                         *endpoint.Covers
	DateFormats                    *endpoint.DateFormats
	EventLogos                     *endpoint.EventLogos
	EventNetworks                  *endpoint.EventNetworks
	Events                         *endpoint.Events
	ExternalGames                  *endpoint.ExternalGames
	ExternalGameSources            *endpoint.ExternalGameSources
	Franchises                     *endpoint.Franchises
	GameEngineLogos                *endpoint.GameEngineLogos
	GameEngines                    *endpoint.GameEngines
	GameLocalizations              *endpoint.GameLocalizations
	GameModes                      *endpoint.GameModes
	GameReleaseFormats             *endpoint.GameReleaseFormats
	Games                          *endpoint.Games
	GameStatuses                   *endpoint.GameStatuses
	GameTimeToBeats                *endpoint.GameTimeToBeats
	GameTypes                      *endpoint.GameTypes
	GameVersionFeatures            *endpoint.GameVersionFeatures
	GameVersionFeatureValues       *endpoint.GameVersionFeatureValues
	GameVersions                   *endpoint.GameVersions
	GameVideos                     *endpoint.GameVideos
	Genres                         *endpoint.Genres
	InvolvedCompanies              *endpoint.InvolvedCompanies
	Keywords                       *endpoint.Keywords
	Languages                      *endpoint.Languages
	LanguageSupports               *endpoint.LanguageSupports
	LanguageSupportTypes           *endpoint.LanguageSupportTypes
	MultiplayerModes               *endpoint.MultiplayerModes
	NetworkTypes                   *endpoint.NetworkTypes
	PlatformFamilies               *endpoint.PlatformFamilies
	PlatformLogos                  *endpoint.PlatformLogos
	Platforms                      *endpoint.Platforms
	PlatformTypes                  *endpoint.PlatformTypes
	PlatformVersionCompanies       *endpoint.PlatformVersionCompanies
	PlatformVersionReleaseDates    *endpoint.PlatformVersionReleaseDates
	PlatformVersions               *endpoint.PlatformVersions
	PlatformWebsites               *endpoint.PlatformWebsites
	PlayerPerspectives             *endpoint.PlayerPerspectives
	PopularityPrimitives           *endpoint.PopularityPrimitives
	PopularityTypes                *endpoint.PopularityTypes
	Regions                        *endpoint.Regions
	ReleaseDateRegions             *endpoint.ReleaseDateRegions
	ReleaseDates                   *endpoint.ReleaseDates
	ReleaseDateStatuses            *endpoint.ReleaseDateStatuses
	Screenshots                    *endpoint.Screenshots
	Search                         *endpoint.Search
	Themes                         *endpoint.Themes
	Webhooks                       *endpoint.Webhooks
	Websites                       *endpoint.Websites
	WebsiteTypes                   *endpoint.WebsiteTypes
}

func New(clientID, clientSecret string) *Client {
	c := &Client{
		clientID:        clientID,
		limiter:         newRateLimiter(4),
		token:           NewTwitchToken(clientID, clientSecret),
		flaresolverr:    nil,
		EntityEndpoints: make(map[endpoint.EndpointName]endpoint.EntityEndpoint),
	}

	registerAllEndpoints(c)

	return c
}

func NewWithFlaresolverr(clientID, clientSecret string, f *flaresolverr.Flaresolverr) *Client {
	c := New(clientID, clientSecret)
	c.flaresolverr = f
	return c
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

func GetItemsPagniated[T any](offset, limit int, f func(string) ([]*T, error)) ([]*T, error) {
	query := fmt.Sprintf("offset %d; limit %d; f *; sort id asc;", offset, limit)
	items, err := f(query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemsLength[T any](f func(string) ([]*T, error)) (uint64, error) {
	query := "fields id; sort id desc; limit 1;"
	items, err := f(query)
	if err != nil {
		return 0, err
	}

	if len(items) == 0 {
		return 0, fmt.Errorf("no results: %s", query)
	}

	type Iid interface {
		GetId() uint64
	}

	item, ok := any(items[0]).(Iid)

	if !ok {
		return 0, fmt.Errorf("failed to convert")
	}

	return item.GetId(), nil
}

func GetItemByID[T any](id uint64, f func(string) ([]*T, error)) (*T, error) {
	query := fmt.Sprintf("where id = %d; fields *;", id)
	items, err := f(query)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return items[0], nil
}

func GetItemsByIDs[T any](ids []uint64, f func(string) ([]*T, error)) ([]*T, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	items, err := f(idStr)
	if err != nil {
		return nil, err
	}

	return items, nil
}
