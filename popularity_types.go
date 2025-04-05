package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPopularityTypes(query string) ([]*pb.PopularityType, error) {
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

func (g *Client) GetPopularityTypeByID(id uint64) (*pb.PopularityType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	popularityTypes, err := g.GetPopularityTypes(query)
	if err != nil {
		return nil, err
	}
	return popularityTypes[0], nil
}

func (g *Client) GetPopularityTypesByIDs(ids []uint64) ([]*pb.PopularityType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPopularityTypes(idStr)
}

func (g *Client) GetPopularityTypesByExternalPopularitySourceID(id uint64) ([]*pb.PopularityType, error) {
	query := fmt.Sprintf(`where external_popularity_source = %d; fields *;`, id)
	return g.GetPopularityTypes(query)
}

func (g *Client) GetPopularityTypesByExternalPopularitySourceIDs(ids []uint64) ([]*pb.PopularityType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where external_popularity_source = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPopularityTypes(idStr)
}

func (g *Client) GetPopularityTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	popularityTypes, err := g.GetPopularityTypes(query)
	if err != nil {
		return 0, err
	}
	return int(popularityTypes[0].Id), nil
}
