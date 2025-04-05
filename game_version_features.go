package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameVersionFeatures(query string) ([]*pb.GameVersionFeature, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_version_features.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionFeatureResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameversionfeatures) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameversionfeatures, nil
}

func (g *Client) GetGameVersionFeatureByID(id uint64) (*pb.GameVersionFeature, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameVersionFeatures, err := g.GetGameVersionFeatures(query)
	if err != nil {
		return nil, err
	}
	return gameVersionFeatures[0], nil
}

func (g *Client) GetGameVersionFeaturesByIDs(ids []uint64) ([]*pb.GameVersionFeature, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersionFeatures(idStr)
}

func (g *Client) GetGameVersionFeaturesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	gameVersionFeatures, err := g.GetGameVersionFeatures(query)
	if err != nil {
		return 0, err
	}
	return int(gameVersionFeatures[0].Id), nil
}
