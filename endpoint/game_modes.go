package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameModes struct {
	BaseEndpoint[pb.GameMode]
}

func NewGameModes(request func(URL string, dataBody any) (*resty.Response, error)) *GameModes {
	a := &GameModes{
		BaseEndpoint[pb.GameMode]{
			endpointName: EPGameModes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameModes) Query(query string) ([]*pb.GameMode, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_modes.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameModeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamemodes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamemodes, nil
}
