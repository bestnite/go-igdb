package igdb

import (
	"fmt"

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
