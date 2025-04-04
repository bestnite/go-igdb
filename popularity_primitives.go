package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPopularityPrimitives(query string) ([]*pb.PopularityPrimitive, error) {
	resp, err := g.Request("https://api.igdb.com/v4/popularity_primitives.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PopularityPrimitiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Popularityprimitives) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Popularityprimitives, nil
}

func (g *igdb) GetPopularityPrimitiveByID(id uint64) (*pb.PopularityPrimitive, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	popularityPrimitives, err := g.GetPopularityPrimitives(query)
	if err != nil {
		return nil, err
	}
	return popularityPrimitives[0], nil
}

func (g *igdb) GetPopularityPrimitivesByIDs(ids []uint64) ([]*pb.PopularityPrimitive, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPopularityPrimitives(idStr)
}

// GetPopularityPrimitive retrieves popular IGDB game IDs based on a given popularity type.
// popularity_type = 1 IGDB Visits
// popularity_type = 2 IGDB Want to Play
// popularity_type = 3 IGDB Playing
// popularity_type = 4 IGDB Played
func (g *igdb) GetPopularityPrimitivesByPopularityType(popularityType, offset, limit int) ([]*pb.PopularityPrimitive, error) {
	query := fmt.Sprintf("fields game_id,value,popularity_type; sort value desc; limit %d; offset %d; where popularity_type = %d;", limit, offset, popularityType)
	return g.GetPopularityPrimitives(query)
}

func (g *igdb) GetPopularityPrimitivesByExternalPopularitySourceID(id uint64) ([]*pb.PopularityPrimitive, error) {
	query := fmt.Sprintf(`where external_popularity_source = %d; fields *;`, id)
	return g.GetPopularityPrimitives(query)
}

func (g *igdb) GetPopularityPrimitivesByExternalPopularitySourceIDs(ids []uint64) ([]*pb.PopularityPrimitive, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where external_popularity_source = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPopularityPrimitives(idStr)
}

func (g *igdb) GetPopularityPrimitivesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	popularityPrimitives, err := g.GetPopularityPrimitives(query)
	if err != nil {
		return 0, err
	}
	return int(popularityPrimitives[0].Id), nil
}
