package igdb

import "github.com/bestnite/go-igdb/endpoint"

func registerAllEndpoints(c *Client) {
	c.AgeRatingCategories = &endpoint.AgeRatingCategories{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAgeRatingCategories),
	}
	c.EntityEndpoints[endpoint.EPAgeRatingCategories] = c.AgeRatingCategories

	c.AgeRatingContentDescriptions = &endpoint.AgeRatingContentDescriptions{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAgeRatingContentDescriptions),
	}
	c.EntityEndpoints[endpoint.EPAgeRatingContentDescriptions] = c.AgeRatingContentDescriptions

	c.AgeRatingContentDescriptionsV2 = &endpoint.AgeRatingContentDescriptionsV2{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAgeRatingContentDescriptionsV2),
	}
	c.EntityEndpoints[endpoint.EPAgeRatingContentDescriptionsV2] = c.AgeRatingContentDescriptionsV2

	c.AgeRatingOrganizations = &endpoint.AgeRatingOrganizations{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAgeRatingOrganizations),
	}
	c.EntityEndpoints[endpoint.EPAgeRatingOrganizations] = c.AgeRatingOrganizations

	c.AgeRatings = &endpoint.AgeRatings{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAgeRatings),
	}
	c.EntityEndpoints[endpoint.EPAgeRatings] = c.AgeRatings

	c.AlternativeNames = &endpoint.AlternativeNames{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPAlternativeNames),
	}
	c.EntityEndpoints[endpoint.EPAlternativeNames] = c.AlternativeNames

	c.Artworks = &endpoint.Artworks{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPArtworks),
	}
	c.EntityEndpoints[endpoint.EPArtworks] = c.Artworks

	c.CharacterGenders = &endpoint.CharacterGenders{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCharacterGenders),
	}
	c.EntityEndpoints[endpoint.EPCharacterGenders] = c.CharacterGenders

	c.CharacterMugShots = &endpoint.CharacterMugShots{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCharacterMugShots),
	}
	c.EntityEndpoints[endpoint.EPCharacterMugShots] = c.CharacterMugShots

	c.Characters = &endpoint.Characters{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCharacters),
	}
	c.EntityEndpoints[endpoint.EPCharacters] = c.Characters

	c.CharacterSpecies = &endpoint.CharacterSpecies{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCharacterSpecies),
	}
	c.EntityEndpoints[endpoint.EPCharacterSpecies] = c.CharacterSpecies

	c.CollectionMemberships = &endpoint.CollectionMemberships{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollectionMemberships),
	}
	c.EntityEndpoints[endpoint.EPCollectionMemberships] = c.CollectionMemberships

	c.CollectionMembershipTypes = &endpoint.CollectionMembershipTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollectionMembershipTypes),
	}
	c.EntityEndpoints[endpoint.EPCollectionMembershipTypes] = c.CollectionMembershipTypes

	c.CollectionRelations = &endpoint.CollectionRelations{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollectionRelations),
	}
	c.EntityEndpoints[endpoint.EPCollectionRelations] = c.CollectionRelations

	c.CollectionRelationTypes = &endpoint.CollectionRelationTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollectionRelationTypes),
	}
	c.EntityEndpoints[endpoint.EPCollectionRelationTypes] = c.CollectionRelationTypes

	c.Collections = &endpoint.Collections{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollections),
	}
	c.EntityEndpoints[endpoint.EPCollections] = c.Collections

	c.CollectionTypes = &endpoint.CollectionTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCollectionTypes),
	}
	c.EntityEndpoints[endpoint.EPCollectionTypes] = c.CollectionTypes

	c.Companies = &endpoint.Companies{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCompanies),
	}
	c.EntityEndpoints[endpoint.EPCompanies] = c.Companies

	c.CompanyLogos = &endpoint.CompanyLogos{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCompanyLogos),
	}
	c.EntityEndpoints[endpoint.EPCompanyLogos] = c.CompanyLogos

	c.CompanyStatuses = &endpoint.CompanyStatuses{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCompanyStatuses),
	}
	c.EntityEndpoints[endpoint.EPCompanyStatuses] = c.CompanyStatuses

	c.CompanyWebsites = &endpoint.CompanyWebsites{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCompanyWebsites),
	}
	c.EntityEndpoints[endpoint.EPCompanyWebsites] = c.CompanyWebsites

	c.Covers = &endpoint.Covers{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPCovers),
	}
	c.EntityEndpoints[endpoint.EPCovers] = c.Covers

	c.DateFormats = &endpoint.DateFormats{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPDateFormats),
	}
	c.EntityEndpoints[endpoint.EPDateFormats] = c.DateFormats

	c.EventLogos = &endpoint.EventLogos{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPEventLogos),
	}
	c.EntityEndpoints[endpoint.EPEventLogos] = c.EventLogos

	c.EventNetworks = &endpoint.EventNetworks{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPEventNetworks),
	}
	c.EntityEndpoints[endpoint.EPEventNetworks] = c.EventNetworks

	c.Events = &endpoint.Events{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPEvents),
	}
	c.EntityEndpoints[endpoint.EPEvents] = c.Events

	c.ExternalGames = &endpoint.ExternalGames{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPExternalGames),
	}
	c.EntityEndpoints[endpoint.EPExternalGames] = c.ExternalGames

	c.ExternalGameSources = &endpoint.ExternalGameSources{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPExternalGameSources),
	}
	c.EntityEndpoints[endpoint.EPExternalGameSources] = c.ExternalGameSources

	c.Franchises = &endpoint.Franchises{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPFranchises),
	}
	c.EntityEndpoints[endpoint.EPFranchises] = c.Franchises

	c.GameEngineLogos = &endpoint.GameEngineLogos{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameEngineLogos),
	}
	c.EntityEndpoints[endpoint.EPGameEngineLogos] = c.GameEngineLogos

	c.GameEngines = &endpoint.GameEngines{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameEngines),
	}
	c.EntityEndpoints[endpoint.EPGameEngines] = c.GameEngines

	c.GameLocalizations = &endpoint.GameLocalizations{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameLocalizations),
	}
	c.EntityEndpoints[endpoint.EPGameLocalizations] = c.GameLocalizations

	c.GameModes = &endpoint.GameModes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameModes),
	}
	c.EntityEndpoints[endpoint.EPGameModes] = c.GameModes

	c.GameReleaseFormats = &endpoint.GameReleaseFormats{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameReleaseFormats),
	}
	c.EntityEndpoints[endpoint.EPGameReleaseFormats] = c.GameReleaseFormats

	c.Games = &endpoint.Games{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGames),
	}
	c.EntityEndpoints[endpoint.EPGames] = c.Games

	c.GameStatuses = &endpoint.GameStatuses{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameStatuses),
	}
	c.EntityEndpoints[endpoint.EPGameStatuses] = c.GameStatuses

	c.GameTimeToBeats = &endpoint.GameTimeToBeats{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameTimeToBeats),
	}
	c.EntityEndpoints[endpoint.EPGameTimeToBeats] = c.GameTimeToBeats

	c.GameTypes = &endpoint.GameTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameTypes),
	}
	c.EntityEndpoints[endpoint.EPGameTypes] = c.GameTypes

	c.GameVersionFeatures = &endpoint.GameVersionFeatures{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameVersionFeatures),
	}
	c.EntityEndpoints[endpoint.EPGameVersionFeatures] = c.GameVersionFeatures

	c.GameVersionFeatureValues = &endpoint.GameVersionFeatureValues{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameVersionFeatureValues),
	}
	c.EntityEndpoints[endpoint.EPGameVersionFeatureValues] = c.GameVersionFeatureValues

	c.GameVersions = &endpoint.GameVersions{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameVersions),
	}
	c.EntityEndpoints[endpoint.EPGameVersions] = c.GameVersions

	c.GameVideos = &endpoint.GameVideos{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGameVideos),
	}
	c.EntityEndpoints[endpoint.EPGameVideos] = c.GameVideos

	c.Genres = &endpoint.Genres{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPGenres),
	}
	c.EntityEndpoints[endpoint.EPGenres] = c.Genres

	c.InvolvedCompanies = &endpoint.InvolvedCompanies{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPInvolvedCompanies),
	}
	c.EntityEndpoints[endpoint.EPInvolvedCompanies] = c.InvolvedCompanies

	c.Keywords = &endpoint.Keywords{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPKeywords),
	}
	c.EntityEndpoints[endpoint.EPKeywords] = c.Keywords

	c.Languages = &endpoint.Languages{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPLanguages),
	}
	c.EntityEndpoints[endpoint.EPLanguages] = c.Languages

	c.LanguageSupports = &endpoint.LanguageSupports{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPLanguageSupports),
	}
	c.EntityEndpoints[endpoint.EPLanguageSupports] = c.LanguageSupports

	c.LanguageSupportTypes = &endpoint.LanguageSupportTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPLanguageSupportTypes),
	}
	c.EntityEndpoints[endpoint.EPLanguageSupportTypes] = c.LanguageSupportTypes

	c.MultiplayerModes = &endpoint.MultiplayerModes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPMultiplayerModes),
	}
	c.EntityEndpoints[endpoint.EPMultiplayerModes] = c.MultiplayerModes

	c.NetworkTypes = &endpoint.NetworkTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPNetworkTypes),
	}
	c.EntityEndpoints[endpoint.EPNetworkTypes] = c.NetworkTypes

	c.PlatformFamilies = &endpoint.PlatformFamilies{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformFamilies),
	}
	c.EntityEndpoints[endpoint.EPPlatformFamilies] = c.PlatformFamilies

	c.PlatformLogos = &endpoint.PlatformLogos{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformLogos),
	}
	c.EntityEndpoints[endpoint.EPPlatformLogos] = c.PlatformLogos

	c.Platforms = &endpoint.Platforms{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatforms),
	}
	c.EntityEndpoints[endpoint.EPPlatforms] = c.Platforms

	c.PlatformTypes = &endpoint.PlatformTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformTypes),
	}
	c.EntityEndpoints[endpoint.EPPlatformTypes] = c.PlatformTypes

	c.PlatformVersionCompanies = &endpoint.PlatformVersionCompanies{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformVersionCompanies),
	}
	c.EntityEndpoints[endpoint.EPPlatformVersionCompanies] = c.PlatformVersionCompanies

	c.PlatformVersionReleaseDates = &endpoint.PlatformVersionReleaseDates{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformVersionReleaseDates),
	}
	c.EntityEndpoints[endpoint.EPPlatformVersionReleaseDates] = c.PlatformVersionReleaseDates

	c.PlatformVersions = &endpoint.PlatformVersions{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformVersions),
	}
	c.EntityEndpoints[endpoint.EPPlatformVersions] = c.PlatformVersions

	c.PlatformWebsites = &endpoint.PlatformWebsites{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlatformWebsites),
	}
	c.EntityEndpoints[endpoint.EPPlatformWebsites] = c.PlatformWebsites

	c.PlayerPerspectives = &endpoint.PlayerPerspectives{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPlayerPerspectives),
	}
	c.EntityEndpoints[endpoint.EPPlayerPerspectives] = c.PlayerPerspectives

	c.PopularityPrimitives = &endpoint.PopularityPrimitives{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPopularityPrimitives),
	}
	c.EntityEndpoints[endpoint.EPPopularityPrimitives] = c.PopularityPrimitives

	c.PopularityTypes = &endpoint.PopularityTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPPopularityTypes),
	}
	c.EntityEndpoints[endpoint.EPPopularityTypes] = c.PopularityTypes

	c.Regions = &endpoint.Regions{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPRegions),
	}
	c.EntityEndpoints[endpoint.EPRegions] = c.Regions

	c.ReleaseDateRegions = &endpoint.ReleaseDateRegions{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPReleaseDateRegions),
	}
	c.EntityEndpoints[endpoint.EPReleaseDateRegions] = c.ReleaseDateRegions

	c.ReleaseDates = &endpoint.ReleaseDates{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPReleaseDates),
	}
	c.EntityEndpoints[endpoint.EPReleaseDates] = c.ReleaseDates

	c.ReleaseDateStatuses = &endpoint.ReleaseDateStatuses{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPReleaseDateStatuses),
	}
	c.EntityEndpoints[endpoint.EPReleaseDateStatuses] = c.ReleaseDateStatuses

	c.Screenshots = &endpoint.Screenshots{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPScreenshots),
	}
	c.EntityEndpoints[endpoint.EPScreenshots] = c.Screenshots

	c.Themes = &endpoint.Themes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPThemes),
	}
	c.EntityEndpoints[endpoint.EPThemes] = c.Themes

	c.Websites = &endpoint.Websites{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPWebsites),
	}
	c.EntityEndpoints[endpoint.EPWebsites] = c.Websites

	c.WebsiteTypes = &endpoint.WebsiteTypes{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPWebsiteTypes),
	}
	c.EntityEndpoints[endpoint.EPWebsiteTypes] = c.WebsiteTypes

	c.Webhooks = &endpoint.Webhooks{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPWebhooks),
	}

	c.Search = &endpoint.Search{
		BaseEndpoint: *endpoint.NewBaseEndpoint(c.Request, endpoint.EPSearch),
	}
}
