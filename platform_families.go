package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformFamilies(query string) ([]*pb.PlatformFamily, error) {
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

func (g *igdb) GetPlatformFamilyByID(id uint64) (*pb.PlatformFamily, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformFamilies, err := g.GetPlatformFamilies(query)
	if err != nil {
		return nil, err
	}
	return platformFamilies[0], nil
}

func (g *igdb) GetPlatformFamiliesByIDs(ids []uint64) ([]*pb.PlatformFamily, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformFamilies(idStr)
}

func (g *igdb) GetPlatformFamiliesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platformFamilies, err := g.GetPlatformFamilies(query)
	if err != nil {
		return 0, err
	}
	return int(platformFamilies[0].Id), nil
}
