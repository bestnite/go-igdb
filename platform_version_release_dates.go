package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPlatformVersionReleaseDates(query string) ([]*pb.PlatformVersionReleaseDate, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_version_release_dates.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversionreleasedates) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversionreleasedates, nil
}
