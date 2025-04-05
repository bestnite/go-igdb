package igdb

import (
	"github.com/bestnite/go-igdb/endpoint"
)

func registerAllEndpoints(c *Client) {
	c.AgeRatingCategories = endpoint.NewAgeRatingCategories(c.Request)

	c.AgeRatingContentDescriptions = endpoint.NewAgeRatingContentDescriptions(c.Request)

	c.AgeRatingContentDescriptionsV2 = endpoint.NewAgeRatingContentDescriptionsV2(c.Request)

	c.AgeRatingOrganizations = endpoint.NewAgeRatingOrganizations(c.Request)

	c.AgeRatings = endpoint.NewAgeRatings(c.Request)

	c.AlternativeNames = endpoint.NewAlternativeNames(c.Request)

	c.Artworks = endpoint.NewArtworks(c.Request)

	c.CharacterGenders = endpoint.NewCharacterGenders(c.Request)

	c.CharacterMugShots = endpoint.NewCharacterMugShots(c.Request)

	c.Characters = endpoint.NewCharacters(c.Request)

	c.CharacterSpecies = endpoint.NewCharacterSpecies(c.Request)

	c.CollectionMemberships = endpoint.NewCollectionMemberships(c.Request)

	c.CollectionMembershipTypes = endpoint.NewCollectionMembershipTypes(c.Request)

	c.CollectionRelations = endpoint.NewCollectionRelations(c.Request)

	c.CollectionRelationTypes = endpoint.NewCollectionRelationTypes(c.Request)

	c.Collections = endpoint.NewCollections(c.Request)

	c.CollectionTypes = endpoint.NewCollectionTypes(c.Request)

	c.Companies = endpoint.NewCompanies(c.Request)

	c.CompanyLogos = endpoint.NewCompanyLogos(c.Request)

	c.CompanyStatuses = endpoint.NewCompanyStatuses(c.Request)

	c.CompanyWebsites = endpoint.NewCompanyWebsites(c.Request)

	c.Covers = endpoint.NewCovers(c.Request)

	c.DateFormats = endpoint.NewDateFormats(c.Request)

	c.EventLogos = endpoint.NewEventLogos(c.Request)

	c.EventNetworks = endpoint.NewEventNetworks(c.Request)

	c.Events = endpoint.NewEvents(c.Request)

	c.ExternalGames = endpoint.NewExternalGames(c.Request)

	c.ExternalGameSources = endpoint.NewExternalGameSources(c.Request)

	c.Franchises = endpoint.NewFranchises(c.Request)

	c.GameEngineLogos = endpoint.NewGameEngineLogos(c.Request)

	c.GameEngines = endpoint.NewGameEngines(c.Request)

	c.GameLocalizations = endpoint.NewGameLocalizations(c.Request)

	c.GameModes = endpoint.NewGameModes(c.Request)

	c.GameReleaseFormats = endpoint.NewGameReleaseFormats(c.Request)

	c.Games = endpoint.NewGames(c.Request)

	c.GameStatuses = endpoint.NewGameStatuses(c.Request)

	c.GameTimeToBeats = endpoint.NewGameTimeToBeats(c.Request)

	c.GameTypes = endpoint.NewGameTypes(c.Request)

	c.GameVersionFeatures = endpoint.NewGameVersionFeatures(c.Request)

	c.GameVersionFeatureValues = endpoint.NewGameVersionFeatureValues(c.Request)

	c.GameVersions = endpoint.NewGameVersions(c.Request)

	c.GameVideos = endpoint.NewGameVideos(c.Request)

	c.Genres = endpoint.NewGenres(c.Request)

	c.InvolvedCompanies = endpoint.NewInvolvedCompanies(c.Request)

	c.Keywords = endpoint.NewKeywords(c.Request)

	c.Languages = endpoint.NewLanguages(c.Request)

	c.LanguageSupports = endpoint.NewLanguageSupports(c.Request)

	c.LanguageSupportTypes = endpoint.NewLanguageSupportTypes(c.Request)

	c.MultiplayerModes = endpoint.NewMultiplayerModes(c.Request)

	c.NetworkTypes = endpoint.NewNetworkTypes(c.Request)

	c.PlatformFamilies = endpoint.NewPlatformFamilies(c.Request)

	c.PlatformLogos = endpoint.NewPlatformLogos(c.Request)

	c.Platforms = endpoint.NewPlatforms(c.Request)

	c.PlatformTypes = endpoint.NewPlatformTypes(c.Request)

	c.PlatformVersionCompanies = endpoint.NewPlatformVersionCompanies(c.Request)

	c.PlatformVersionReleaseDates = endpoint.NewPlatformVersionReleaseDates(c.Request)

	c.PlatformVersions = endpoint.NewPlatformVersions(c.Request)

	c.PlatformWebsites = endpoint.NewPlatformWebsites(c.Request)

	c.PlayerPerspectives = endpoint.NewPlayerPerspectives(c.Request)

	c.PopularityPrimitives = endpoint.NewPopularityPrimitives(c.Request)

	c.PopularityTypes = endpoint.NewPopularityTypes(c.Request)

	c.Regions = endpoint.NewRegions(c.Request)

	c.ReleaseDateRegions = endpoint.NewReleaseDateRegions(c.Request)

	c.ReleaseDates = endpoint.NewReleaseDates(c.Request)

	c.ReleaseDateStatuses = endpoint.NewReleaseDateStatuses(c.Request)

	c.Screenshots = endpoint.NewScreenshots(c.Request)

	c.Themes = endpoint.NewThemes(c.Request)

	c.Websites = endpoint.NewWebsites(c.Request)

	c.WebsiteTypes = endpoint.NewWebsiteTypes(c.Request)

	c.Webhooks = endpoint.NewWebhooks(c.Request)

	c.Search = endpoint.NewSearch(c.Request)
}
