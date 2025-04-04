package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPopularityTypes(query string) ([]*pb.PopularityType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/popularity_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PopularityTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Popularitytypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Popularitytypes, nil
}

func (g *igdb) GetPopularityTypeByID(id uint64) (*pb.PopularityType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	popularityTypes, err := g.GetPopularityTypes(query)
	if err != nil {
		return nil, err
	}
	return popularityTypes[0], nil
}

func (g *igdb) GetPopularityTypesByIDs(ids []uint64) ([]*pb.PopularityType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPopularityTypes(idStr)
}
