package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetReleaseDateRegions(query string) ([]*pb.ReleaseDateRegion, error) {
	resp, err := g.Request("https://api.igdb.com/v4/release_date_regions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateRegionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedateregions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedateregions, nil
}

func (g *Client) GetReleaseDateRegionByID(id uint64) (*pb.ReleaseDateRegion, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	releaseDateRegions, err := g.GetReleaseDateRegions(query)
	if err != nil {
		return nil, err
	}
	return releaseDateRegions[0], nil
}

func (g *Client) GetReleaseDateRegionsByIDs(ids []uint64) ([]*pb.ReleaseDateRegion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetReleaseDateRegions(idStr)
}

func (g *Client) GetReleaseDateRegionsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	releaseDateRegions, err := g.GetReleaseDateRegions(query)
	if err != nil {
		return 0, err
	}
	return int(releaseDateRegions[0].Id), nil
}
