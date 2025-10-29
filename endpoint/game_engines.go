package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameEngines struct {
	BaseEndpoint[pb.GameEngine]
}

func NewGameEngines(request RequestFunc) *GameEngines {
	a := &GameEngines{
		BaseEndpoint: BaseEndpoint[pb.GameEngine]{
			endpointName: EPGameEngines,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameEngines) Query(ctx context.Context, query string) ([]*pb.GameEngine, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameEngineResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Gameengines, nil
}
