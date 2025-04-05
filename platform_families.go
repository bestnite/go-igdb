package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPlatformFamilies(query string) ([]*pb.PlatformFamily, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_families.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformFamilyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformfamilies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformfamilies, nil
}
