package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatforms(query string) ([]*pb.Platform, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platforms.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platforms) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}
	return data.Platforms, nil
}

func (g *igdb) GetPlatformByID(id uint64) (*pb.Platform, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platforms, err := g.GetPlatforms(query)
	if err != nil {
		return nil, err
	}
	return platforms[0], nil
}

func (g *igdb) GetPlatformsByIDs(ids []uint64) ([]*pb.Platform, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatforms(idStr)
}
