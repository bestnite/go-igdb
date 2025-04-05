package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPlatformVersions(query string) ([]*pb.PlatformVersion, error) {
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

func (g *Client) GetPlatformVersionByID(id uint64) (*pb.PlatformVersion, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformVersions, err := g.GetPlatformVersions(query)
	if err != nil {
		return nil, err
	}
	return platformVersions[0], nil
}

func (g *Client) GetPlatformVersionsByIDs(ids []uint64) ([]*pb.PlatformVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersions(idStr)
}

func (g *Client) GetPlatformVersionsByMainManufacturerID(id uint64) ([]*pb.PlatformVersion, error) {
	query := fmt.Sprintf(`where main_manufacturer = %d; fields *;`, id)
	return g.GetPlatformVersions(query)
}

func (g *Client) GetPlatformVersionsByMainManufacturerIDs(ids []uint64) ([]*pb.PlatformVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where main_manufacturer = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersions(idStr)
}

func (g *Client) GetPlatformVersionsByPlatformLogoID(id uint64) ([]*pb.PlatformVersion, error) {
	query := fmt.Sprintf(`where platform_logo = %d; fields *;`, id)
	return g.GetPlatformVersions(query)
}

func (g *Client) GetPlatformVersionsByPlatformLogoIDs(ids []uint64) ([]*pb.PlatformVersion, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersions(idStr)
}

func (g *Client) GetPlatformVersionsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platformVersions, err := g.GetPlatformVersions(query)
	if err != nil {
		return 0, err
	}
	return int(platformVersions[0].Id), nil
}
