package igdb

import (
	"fmt"
	"github/bestnite/go-igdb/constant"
	"strconv"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGame(id uint64) (*pb.Game, error) {
	resp, err := g.Request(constant.IGDBGameURL, fmt.Sprintf(`where id = %v; fields *, age_ratings.*, alternative_names.*, artworks.*, collection.*, cover.*, external_games.*, external_games.platform.* , franchise.*, game_engines.*, game_engines.logo.*, game_engines.companies.* , game_modes.*, genres.*, involved_companies.*, involved_companies.company.* , keywords.*, multiplayer_modes.*, multiplayer_modes.platform.*, platforms.*, platforms.platform_logo.*, platforms.platform_family.*, platforms.versions.*, platforms.websites.* , player_perspectives.*, release_dates.*, release_dates.platform.*, release_dates.status.* , screenshots.*, themes.*, videos.*, websites.*, language_supports.*, language_supports.language.*, language_supports.language_support_type.* , game_localizations.*, game_localizations.region.* , collections.*, collections.type.*, collections.as_parent_relations.child_collection.*, collections.as_parent_relations.parent_collection.*, collections.as_parent_relations.type.*,collections.as_child_relations.child_collection.*, collections.as_child_relations.parent_collection.*, collections.as_child_relations.type.*, age_ratings.content_descriptions.*, cover.game_localization.*;`, id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch game detail for ID %d: %w", id, err)
	}

	res := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal game detail response: %w", err)
	}

	if len(res.Games) == 0 {
		return nil, fmt.Errorf("failed to fetch game detail for ID %d", id)
	}

	if res.Games[0].Name == "" {
		return g.GetGame(id)
	}

	return res.Games[0], nil
}

func (g *igdb) GetGames(ids []uint64) ([]*pb.Game, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = strconv.FormatUint(id, 10)
	}

	idStr := strings.Join(idStrSlice, ",")

	resp, err := g.Request(constant.IGDBGameURL, fmt.Sprintf(`where id = (%s); fields *, age_ratings.*, alternative_names.*, artworks.*, collection.*, cover.*, external_games.*, external_games.platform.* , franchise.*, game_engines.*, game_engines.logo.*, game_engines.companies.* , game_modes.*, genres.*, involved_companies.*, involved_companies.company.* , keywords.*, multiplayer_modes.*, multiplayer_modes.platform.*, platforms.*, platforms.platform_logo.*, platforms.platform_family.*, platforms.versions.*, platforms.websites.* , player_perspectives.*, release_dates.*, release_dates.platform.*, release_dates.status.* , screenshots.*, themes.*, videos.*, websites.*, language_supports.*, language_supports.language.*, language_supports.language_support_type.* , game_localizations.*, game_localizations.region.* , collections.*, collections.type.*, collections.as_parent_relations.child_collection.*, collections.as_parent_relations.parent_collection.*, collections.as_parent_relations.type.*,collections.as_child_relations.child_collection.*, collections.as_child_relations.parent_collection.*, collections.as_child_relations.type.*, age_ratings.content_descriptions.*, cover.game_localization.*;`, idStr))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IGDB games detail for ID %s: %w", idStr, err)
	}

	res := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal IGDB games detail response: %w", err)
	}

	if len(res.Games) == 0 {
		return nil, fmt.Errorf("failed to fetch IGDB games detail for ID %s", idStr)
	}

	if res.Games[0].Name == "" {
		return g.GetGames(ids)
	}

	return res.Games, nil
}
