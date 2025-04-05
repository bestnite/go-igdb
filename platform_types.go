package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPlatformTypes(query string) ([]*pb.PlatformType, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_types.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformtypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformtypes, nil
}
