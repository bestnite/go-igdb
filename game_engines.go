package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGameEngines(query string) ([]*pb.GameEngine, error) {
	resp, err := g.Request("https://api.igdb.com/v4/game_engines.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameengines) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameengines, nil
}
