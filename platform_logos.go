package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformLogos(query string) ([]*pb.PlatformLogo, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_logos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformlogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformlogos, nil
}

func (g *igdb) GetPlatformLogoByID(id uint64) (*pb.PlatformLogo, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformLogos, err := g.GetPlatformLogos(query)
	if err != nil {
		return nil, err
	}
	return platformLogos[0], nil
}

func (g *igdb) GetPlatformLogosByIDs(ids []uint64) ([]*pb.PlatformLogo, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformLogos(idStr)
}

func (g *igdb) GetPlatformLogosLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platformLogos, err := g.GetPlatformLogos(query)
	if err != nil {
		return 0, err
	}
	return int(platformLogos[0].Id), nil
}
