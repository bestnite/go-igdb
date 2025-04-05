package endpoint

type EndpointName string

var (
	EPAgeRatingCategories            EndpointName = "age_rating_categories"
	EPAgeRatingContentDescriptions   EndpointName = "age_rating_content_descriptions"
	EPAgeRatingContentDescriptionsV2 EndpointName = "age_rating_content_descriptions_v2"
	EPAgeRatingOrganizations         EndpointName = "age_rating_organizations"
	EPAgeRatings                     EndpointName = "age_ratings"
	EPAlternativeNames               EndpointName = "alternative_names"
	EPArtworks                       EndpointName = "artworks"
	EPCharacterGenders               EndpointName = "character_genders"
	EPCharacterMugShots              EndpointName = "character_mug_shots"
	EPCharacters                     EndpointName = "characters"
	EPCharacterSpecies               EndpointName = "character_species"
	EPCollectionMemberships          EndpointName = "collection_memberships"
	EPCollectionMembershipTypes      EndpointName = "collection_membership_types"
	EPCollectionRelations            EndpointName = "collection_relations"
	EPCollectionRelationTypes        EndpointName = "collection_relation_types"
	EPCollections                    EndpointName = "collections"
	EPCollectionTypes                EndpointName = "collection_types"
	EPCompanies                      EndpointName = "companies"
	EPCompanyLogos                   EndpointName = "company_logos"
	EPCompanyStatuses                EndpointName = "company_statuses"
	EPCompanyWebsites                EndpointName = "company_websites"
	EPCovers                         EndpointName = "covers"
	EPDateFormats                    EndpointName = "date_formats"
	EPEventLogos                     EndpointName = "event_logos"
	EPEventNetworks                  EndpointName = "event_networks"
	EPEvents                         EndpointName = "events"
	EPExternalGames                  EndpointName = "external_games"
	EPExternalGameSources            EndpointName = "external_game_sources"
	EPFranchises                     EndpointName = "franchises"
	EPGameEngineLogos                EndpointName = "game_engine_logos"
	EPGameEngines                    EndpointName = "game_engines"
	EPGameLocalizations              EndpointName = "game_localizations"
	EPGameModes                      EndpointName = "game_modes"
	EPGameReleaseFormats             EndpointName = "game_release_formats"
	EPGames                          EndpointName = "games"
	EPGameStatuses                   EndpointName = "game_statuses"
	EPGameTimeToBeats                EndpointName = "game_time_to_beats"
	EPGameTypes                      EndpointName = "game_types"
	EPGameVersionFeatures            EndpointName = "game_version_features"
	EPGameVersionFeatureValues       EndpointName = "game_version_feature_values"
	EPGameVersions                   EndpointName = "game_versions"
	EPGameVideos                     EndpointName = "game_videos"
	EPGenres                         EndpointName = "genres"
	EPInvolvedCompanies              EndpointName = "involved_companies"
	EPKeywords                       EndpointName = "keywords"
	EPLanguages                      EndpointName = "languages"
	EPLanguageSupports               EndpointName = "language_supports"
	EPLanguageSupportTypes           EndpointName = "language_support_types"
	EPMultiplayerModes               EndpointName = "multiplayer_modes"
	EPNetworkTypes                   EndpointName = "network_types"
	EPPlatformFamilies               EndpointName = "platform_families"
	EPPlatformLogos                  EndpointName = "platform_logos"
	EPPlatforms                      EndpointName = "platforms"
	EPPlatformTypes                  EndpointName = "platform_types"
	EPPlatformVersionCompanies       EndpointName = "platform_version_companies"
	EPPlatformVersionReleaseDates    EndpointName = "platform_version_release_dates"
	EPPlatformVersions               EndpointName = "platform_versions"
	EPPlatformWebsites               EndpointName = "platform_websites"
	EPPlayerPerspectives             EndpointName = "player_perspectives"
	EPPopularityPrimitives           EndpointName = "popularity_primitives"
	EPPopularityTypes                EndpointName = "popularity_types"
	EPRegions                        EndpointName = "regions"
	EPReleaseDateRegions             EndpointName = "release_date_regions"
	EPReleaseDates                   EndpointName = "release_dates"
	EPReleaseDateStatuses            EndpointName = "release_date_statuses"
	EPScreenshots                    EndpointName = "screenshots"
	EPSearch                         EndpointName = "search"
	EPThemes                         EndpointName = "themes"
	EPWebhooks                       EndpointName = "webhooks"
	EPWebsites                       EndpointName = "websites"
	EPWebsiteTypes                   EndpointName = "website_types"
)

var AllEndpoints = []EndpointName{
	EPAgeRatingCategories,
	EPAgeRatingContentDescriptions,
	EPAgeRatingContentDescriptionsV2,
	EPAgeRatingOrganizations,
	EPAgeRatings,
	EPAlternativeNames,
	EPArtworks,
	EPCharacterGenders,
	EPCharacterMugShots,
	EPCharacters,
	EPCharacterSpecies,
	EPCollectionMemberships,
	EPCollectionMembershipTypes,
	EPCollectionRelations,
	EPCollectionRelationTypes,
	EPCollections,
	EPCollectionTypes,
	EPCompanies,
	EPCompanyLogos,
	EPCompanyStatuses,
	EPCompanyWebsites,
	EPCovers,
	EPDateFormats,
	EPEventLogos,
	EPEventNetworks,
	EPEvents,
	EPExternalGames,
	EPExternalGameSources,
	EPFranchises,
	EPGameEngineLogos,
	EPGameEngines,
	EPGameLocalizations,
	EPGameModes,
	EPGameReleaseFormats,
	EPGames,
	EPGameStatuses,
	EPGameTimeToBeats,
	EPGameTypes,
	EPGameVersionFeatures,
	EPGameVersionFeatureValues,
	EPGameVersions,
	EPGameVideos,
	EPGenres,
	EPInvolvedCompanies,
	EPKeywords,
	EPLanguages,
	EPLanguageSupports,
	EPLanguageSupportTypes,
	EPMultiplayerModes,
	EPNetworkTypes,
	EPPlatformFamilies,
	EPPlatformLogos,
	EPPlatforms,
	EPPlatformTypes,
	EPPlatformVersionCompanies,
	EPPlatformVersionReleaseDates,
	EPPlatformVersions,
	EPPlatformWebsites,
	EPPlayerPerspectives,
	EPPopularityPrimitives,
	EPPopularityTypes,
	EPRegions,
	EPReleaseDateRegions,
	EPReleaseDates,
	EPReleaseDateStatuses,
	EPScreenshots,
	EPSearch,
	EPThemes,
	EPWebhooks,
	EPWebsites,
	EPWebsiteTypes,
}
