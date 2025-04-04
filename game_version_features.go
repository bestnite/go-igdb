package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGameVersionFeatures(query string) ([]*pb.GameVersionFeature, error) {
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

func (g *igdb) GetGameVersionFeatureByID(id uint64) (*pb.GameVersionFeature, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	gameVersionFeatures, err := g.GetGameVersionFeatures(query)
	if err != nil {
		return nil, err
	}
	return gameVersionFeatures[0], nil
}

func (g *igdb) GetGameVersionFeaturesByIDs(ids []uint64) ([]*pb.GameVersionFeature, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGameVersionFeatures(idStr)
}
