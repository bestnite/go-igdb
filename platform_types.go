package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformTypes(query string) ([]*pb.PlatformType, error) {
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

func (g *igdb) GetPlatformTypeByID(id uint64) (*pb.PlatformType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformTypes, err := g.GetPlatformTypes(query)
	if err != nil {
		return nil, err
	}
	return platformTypes[0], nil
}

func (g *igdb) GetPlatformTypesByIDs(ids []uint64) ([]*pb.PlatformType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformTypes(idStr)
}
