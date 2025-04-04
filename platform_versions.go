package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformVersions(query string) ([]*pb.PlatformVersion, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_versions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversions, nil
}

func (g *igdb) GetPlatformVersionByID(id uint64) (*pb.PlatformVersion, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformVersions, err := g.GetPlatformVersions(query)
	if err != nil {
		return nil, err
	}
	return platformVersions[0], nil
}

func (g *igdb) GetPlatformVersionsByIDs(ids []uint64) ([]*pb.PlatformVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersions(idStr)
}
