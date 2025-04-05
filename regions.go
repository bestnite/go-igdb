package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetRegions(query string) ([]*pb.Region, error) {
	resp, err := g.Request("https://api.igdb.com/v4/regions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.RegionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Regions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Regions, nil
}

func (g *Client) GetRegionByID(id uint64) (*pb.Region, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	regions, err := g.GetRegions(query)
	if err != nil {
		return nil, err
	}
	return regions[0], nil
}

func (g *Client) GetRegionsByIDs(ids []uint64) ([]*pb.Region, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetRegions(idStr)
}

func (g *Client) GetRegionsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	regions, err := g.GetRegions(query)
	if err != nil {
		return 0, err
	}
	return int(regions[0].Id), nil
}
