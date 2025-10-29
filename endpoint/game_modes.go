package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameModes struct {
	BaseEndpoint[pb.GameMode]
}

func NewGameModes(request RequestFunc) *GameModes {
	a := &GameModes{
		BaseEndpoint: BaseEndpoint[pb.GameMode]{
			endpointName: EPGameModes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameModes) Query(ctx context.Context, query string) ([]*pb.GameMode, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameModeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Gamemodes, nil
}
