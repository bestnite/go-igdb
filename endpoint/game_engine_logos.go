package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameEngineLogos struct{ BaseEndpoint }

func (a *GameEngineLogos) Query(query string) ([]*pb.GameEngineLogo, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_engine_logos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameenginelogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameenginelogos, nil
}
