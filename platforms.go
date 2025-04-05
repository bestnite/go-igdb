package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetPlatforms(query string) ([]*pb.Platform, error) {
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

func (g *Client) GetPlatformByID(id uint64) (*pb.Platform, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platforms, err := g.GetPlatforms(query)
	if err != nil {
		return nil, err
	}
	return platforms[0], nil
}

func (g *Client) GetPlatformsByIDs(ids []uint64) ([]*pb.Platform, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatforms(idStr)
}

func (g *Client) GetPlatformsByPlatformFamilyID(id uint64) ([]*pb.Platform, error) {
	query := fmt.Sprintf(`where platform_family = %d; fields *;`, id)
	return g.GetPlatforms(query)
}

func (g *Client) GetPlatformsByPlatformFamilyIDs(ids []uint64) ([]*pb.Platform, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_family = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatforms(idStr)
}

func (g *Client) GetPlatformsByPlatformLogoID(id uint64) ([]*pb.Platform, error) {
	query := fmt.Sprintf(`where platform_logo = %d; fields *;`, id)
	return g.GetPlatforms(query)
}

func (g *Client) GetPlatformsByPlatformLogoIDs(ids []uint64) ([]*pb.Platform, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatforms(idStr)
}

func (g *Client) GetPlatformsByPlatformTypeID(id uint64) ([]*pb.Platform, error) {
	query := fmt.Sprintf(`where platform_type = %d; fields *;`, id)
	return g.GetPlatforms(query)
}

func (g *Client) GetPlatformsByPlatformTypeIDs(ids []uint64) ([]*pb.Platform, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform_type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatforms(idStr)
}

func (g *Client) GetPlatformsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	platforms, err := g.GetPlatforms(query)
	if err != nil {
		return 0, err
	}
	return int(platforms[0].Id), nil
}
