package endpoint

type Endpoint string

var (
	AgeRatingCategories            Endpoint = "age_rating_categories"
	AgeRatingContentDescriptions   Endpoint = "age_rating_content_descriptions"
	AgeRatingContentDescriptionsV2 Endpoint = "age_rating_content_descriptions_v2"
	AgeRatingOrganizations         Endpoint = "age_rating_organizations"
	AgeRatings                     Endpoint = "age_ratings"
	AlternativeNames               Endpoint = "alternative_names"
	Artworks                       Endpoint = "artworks"
	CharacterGenders               Endpoint = "character_genders"
	CharacterMugShots              Endpoint = "character_mug_shots"
	Characters                     Endpoint = "characters"
	CharacterSpecies               Endpoint = "character_species"
	CollectionMemberships          Endpoint = "collection_memberships"
	CollectionMembershipTypes      Endpoint = "collection_membership_types"
	CollectionRelations            Endpoint = "collection_relations"
	CollectionRelationTypes        Endpoint = "collection_relation_types"
	Collections                    Endpoint = "collections"
	CollectionTypes                Endpoint = "collection_types"
	Companies                      Endpoint = "companies"
	CompanyLogos                   Endpoint = "company_logos"
	CompanyStatuses                Endpoint = "company_statuses"
	CompanyWebsites                Endpoint = "company_websites"
	Covers                         Endpoint = "covers"
	DateFormats                    Endpoint = "date_formats"
	EventLogos                     Endpoint = "event_logos"
	EventNetworks                  Endpoint = "event_networks"
	Events                         Endpoint = "events"
	ExternalGames                  Endpoint = "external_games"
	ExternalGameSources            Endpoint = "external_game_sources"
	Franchises                     Endpoint = "franchises"
	GameEngineLogos                Endpoint = "game_engine_logos"
	GameEngines                    Endpoint = "game_engines"
	GameLocalizations              Endpoint = "game_localizations"
	GameModes                      Endpoint = "game_modes"
	GameReleaseFormats             Endpoint = "game_release_formats"
	Games                          Endpoint = "games"
	GameStatuses                   Endpoint = "game_statuses"
	GameTimeToBeats                Endpoint = "game_time_to_beats"
	GameTypes                      Endpoint = "game_types"
	GameVersionFeatures            Endpoint = "game_version_features"
	GameVersionFeatureValues       Endpoint = "game_version_feature_values"
	GameVersions                   Endpoint = "game_versions"
	GameVideos                     Endpoint = "game_videos"
	Genres                         Endpoint = "genres"
	InvolvedCompanies              Endpoint = "involved_companies"
	Keywords                       Endpoint = "keywords"
	Languages                      Endpoint = "languages"
	LanguageSupports               Endpoint = "language_supports"
	LanguageSupportTypes           Endpoint = "language_support_types"
	MultiplayerModes               Endpoint = "multiplayer_modes"
	NetworkTypes                   Endpoint = "network_types"
	PlatformFamilies               Endpoint = "platform_families"
	PlatformLogos                  Endpoint = "platform_logos"
	Platforms                      Endpoint = "platforms"
	PlatformTypes                  Endpoint = "platform_types"
	PlatformVersionCompanies       Endpoint = "platform_version_companies"
	PlatformVersionReleaseDates    Endpoint = "platform_version_release_dates"
	PlatformVersions               Endpoint = "platform_versions"
	PlatformWebsites               Endpoint = "platform_websites"
	PlayerPerspectives             Endpoint = "player_perspectives"
	PopularityPrimitives           Endpoint = "popularity_primitives"
	PopularityTypes                Endpoint = "popularity_types"
	Regions                        Endpoint = "regions"
	ReleaseDateRegions             Endpoint = "release_date_regions"
	ReleaseDates                   Endpoint = "release_dates"
	ReleaseDateStatuses            Endpoint = "release_date_statuses"
	Screenshots                    Endpoint = "screenshots"
	Search                         Endpoint = "search"
	Themes                         Endpoint = "themes"
	Webhooks                       Endpoint = "webhooks"
	Websites                       Endpoint = "websites"
	WebsiteTypes                   Endpoint = "website_types"
)

var AllEndpoints = []Endpoint{
	AgeRatingCategories,
	AgeRatingContentDescriptions,
	AgeRatingContentDescriptionsV2,
	AgeRatingOrganizations,
	AgeRatings,
	AlternativeNames,
	Artworks,
	CharacterGenders,
	CharacterMugShots,
	Characters,
	CharacterSpecies,
	CollectionMemberships,
	CollectionMembershipTypes,
	CollectionRelations,
	CollectionRelationTypes,
	Collections,
	CollectionTypes,
	Companies,
	CompanyLogos,
	CompanyStatuses,
	CompanyWebsites,
	Covers,
	DateFormats,
	EventLogos,
	EventNetworks,
	Events,
	ExternalGames,
	ExternalGameSources,
	Franchises,
	GameEngineLogos,
	GameEngines,
	GameLocalizations,
	GameModes,
	GameReleaseFormats,
	Games,
	GameStatuses,
	GameTimeToBeats,
	GameTypes,
	GameVersionFeatures,
	GameVersionFeatureValues,
	GameVersions,
	GameVideos,
	Genres,
	InvolvedCompanies,
	Keywords,
	Languages,
	LanguageSupports,
	LanguageSupportTypes,
	MultiplayerModes,
	NetworkTypes,
	PlatformFamilies,
	PlatformLogos,
	Platforms,
	PlatformTypes,
	PlatformVersionCompanies,
	PlatformVersionReleaseDates,
	PlatformVersions,
	PlatformWebsites,
	PlayerPerspectives,
	PopularityPrimitives,
	PopularityTypes,
	Regions,
	ReleaseDateRegions,
	ReleaseDates,
	ReleaseDateStatuses,
	Screenshots,
	Search,
	Themes,
	Webhooks,
	Websites,
	WebsiteTypes,
}
