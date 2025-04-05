package igdb

import (
	"fmt"
	"strings"

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

func (g *Client) GetPlatformTypeByID(id uint64) (*pb.PlatformType, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformTypes, err := g.GetPlatformTypes(query)
	if err != nil {
		return nil, err
	}
	return platformTypes[0], nil
}

func (g *Client) GetPlatformTypesByIDs(ids []uint64) ([]*pb.PlatformType, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformTypes(idStr)
}

func (g *Client) GetPlatformTypesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platformTypes, err := g.GetPlatformTypes(query)
	if err != nil {
		return 0, err
	}
	return int(platformTypes[0].Id), nil
}
