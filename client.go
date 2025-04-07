package igdb

import (
	"fmt"

	"github.com/bestnite/go-flaresolverr"
	"github.com/bestnite/go-igdb/endpoint"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	clientID     string
	token        *TwitchToken
	flaresolverr *flaresolverr.Flaresolverr
	limiter      *rateLimiter

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
		clientID:     clientID,
		limiter:      newRateLimiter(4),
		token:        NewTwitchToken(clientID, clientSecret),
		flaresolverr: nil,
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
