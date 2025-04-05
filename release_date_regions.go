package igdb

import (
	"fmt"

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
