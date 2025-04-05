package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformWebsites(query string) ([]*pb.PlatformWebsite, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_websites.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformwebsites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformwebsites, nil
}

func (g *igdb) GetPlatformWebsiteByID(id uint64) (*pb.PlatformWebsite, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformWebsites, err := g.GetPlatformWebsites(query)
	if err != nil {
		return nil, err
	}
	return platformWebsites[0], nil
}

func (g *igdb) GetPlatformWebsitesByIDs(ids []uint64) ([]*pb.PlatformWebsite, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformWebsites(idStr)
}

func (g *igdb) GetPlatformWebsitesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platformWebsites, err := g.GetPlatformWebsites(query)
	if err != nil {
		return 0, err
	}
	return int(platformWebsites[0].Id), nil
}
