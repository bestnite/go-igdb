package igdb

import (
	"fmt"
	"github/bestnite/go-igdb/constant"
	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

// GetPopularGameIDs retrieves popular IGDB game IDs based on a given popularity type.
// popularity_type = 1 IGDB Visits: Game page visits on IGDB.com.
// popularity_type = 2 IGDB Want to Play: Additions to IGDB.com users’ “Want to Play” lists.
// popularity_type = 3 IGDB Playing: Additions to IGDB.com users’ “Playing” lists.
// popularity_type = 4 IGDB Played: Additions to IGDB.com users’ “Played” lists.
func (g *igdb) GetPopularGameIDs(popularityType, offset, limit int) ([]uint64, error) {
	query := fmt.Sprintf("fields game_id,value,popularity_type; sort value desc; limit %d; offset %d; where popularity_type = %d;", limit, offset, popularityType)
	resp, err := g.Request(constant.IGDBPopularityURL, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch popular IGDB game IDs for type %d: %w", popularityType, err)
	}
	data := pb.PopularityPrimitiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal IGDB popular games response: %w", err)
	}

	gameIDs := make([]uint64, 0, len(data.Popularityprimitives))
	for _, game := range data.Popularityprimitives {
		gameIDs = append(gameIDs, uint64(game.GameId))
	}

	return gameIDs, nil
}
