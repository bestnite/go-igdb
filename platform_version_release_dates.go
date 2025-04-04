package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformVersionReleaseDates(query string) ([]*pb.PlatformVersionReleaseDate, error) {
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

func (g *igdb) GetPlatformVersionReleaseDateByID(id uint64) (*pb.PlatformVersionReleaseDate, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformVersionReleaseDates, err := g.GetPlatformVersionReleaseDates(query)
	if err != nil {
		return nil, err
	}
	return platformVersionReleaseDates[0], nil
}

func (g *igdb) GetPlatformVersionReleaseDatesByIDs(ids []uint64) ([]*pb.PlatformVersionReleaseDate, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersionReleaseDates(idStr)
}
