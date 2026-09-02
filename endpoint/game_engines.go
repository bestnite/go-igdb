package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameEngines struct {
	BaseEndpoint[pb.GameEngine]
}

func NewGameEngines(request RequestFunc) *GameEngines {
	a := &GameEngines{
		BaseEndpoint[pb.GameEngine]{
			endpointName: EPGameEngines,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameEngineResult) []*pb.GameEngine { return r.Gameengines })
	return a
}
